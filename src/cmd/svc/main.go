// @title Lens Platform Users Microservice
// @version 1.0.0
// @description Go microservice for users, teams, and groups logic.

// @contact.name Yoan Yomba
// @contact.url https://github.com/LensPlatform/Lens

// @license.name MIT License
// @license.url https://github.com/stefanprodan/podinfo/blob/master/LICENSE

// @host localhost:8085
// @BasePath /
// @schemes http https
package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/tabwriter"

	"github.com/alexflint/go-arg"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/lightstep/lightstep-tracer-go"
	"github.com/oklog/oklog/pkg/group"
	stdopentracing "github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"sourcegraph.com/sourcegraph/appdash"
	appdashot "sourcegraph.com/sourcegraph/appdash/opentracing"

	"github.com/LensPlatform/Lens/src/internal/log"
	"github.com/LensPlatform/Lens/src/pkg/config"
	"github.com/LensPlatform/Lens/src/pkg/endpoint"
	models "github.com/LensPlatform/Lens/src/pkg/models/proto"
	"github.com/LensPlatform/Lens/src/pkg/service"
	"github.com/LensPlatform/Lens/src/pkg/transport"
)

func main() {
	// Load config file
	config.DefaultConfiguration()
	arg.MustParse(config.Config)

	// configure logging
	zapLogger, _ := log.InitZap(viper.GetString("level"), config.Config)
	defer zapLogger.Sync()
	stdLog := zap.RedirectStdLog(zapLogger)
	defer stdLog()

	var zipkinTracer *zipkin.Tracer
	{
		if config.Config.ZipkinUrl != "" {
			var (
				err         error
				hostPort    = "8085"
				serviceName = config.Config.Name
				reporter    = zipkinhttp.NewReporter(config.Config.ZipkinUrl)
			)
			defer reporter.Close()
			zEP, err := zipkin.NewEndpoint(serviceName, hostPort)
			if err != nil {
				zapLogger.Error(err.Error())
				os.Exit(1)
			}

			sampler, err := zipkin.NewCountingSampler(1)
			if err != nil {
				zapLogger.Error(err.Error())
				os.Exit(1)
			}

			zipkinTracer, err = zipkin.NewTracer(reporter, zipkin.WithSampler(sampler), zipkin.WithLocalEndpoint(zEP))
			if err != nil {
				zapLogger.Error(err.Error())
				os.Exit(1)
			}
			if !(config.Config.ZipkinBridge) {
				zapLogger.Info("Tracer", zap.String("type of tracer", "zipkin"),
					zap.String("URL", config.Config.ZipkinUrl))
			}
		}
	}

	// Determine which OpenTracing tracer to use. We'll pass the tracer to all the
	// components that use it, as a dependency.
	var tracer stdopentracing.Tracer
	{
		if config.Config.ZipkinBridge && zipkinTracer != nil {
			zapLogger.Info("Tracer", zap.String("type of tracer", "zipkin"),
				zap.String("URL", config.Config.ZipkinUrl))
			tracer = zipkinot.Wrap(zipkinTracer)
			zipkinTracer = nil // do not instrument with both native tracer and opentracing bridge
		} else if config.Config.LightstepToken != "" {
			zapLogger.Info("Tracer", zap.String("type of tracer", "LightStep"))
			tracer = lightstep.NewTracer(lightstep.Options{
				AccessToken: config.Config.LightstepToken,
			})
			defer lightstep.FlushLightStepTracer(tracer)
		} else if config.Config.Appdash != "" {
			zapLogger.Info("Tracer", zap.String("type of tracer", "Appdash"),
				zap.String("Appdash", config.Config.Appdash))
			tracer = appdashot.NewTracer(appdash.NewRemoteCollector(config.Config.Appdash))
		} else {
			tracer = stdopentracing.GlobalTracer() // no-op
		}
	}


	counters := InitMetrics()
	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())

	// configure sql db connection
	db, err := InitDbConnection(zapLogger)
	if err != nil {
		zapLogger.Error(err.Error(), zap.String("Connection Error", "Unable To Connect To Database"))
	}
	defer db.Close()

	// connect to rabbitmq
	amqpproducerconn, amqpconsumerconn := InitQueues(zapLogger)

	httpHandler := InitService(zapLogger, db, amqpproducerconn, amqpconsumerconn, counters, tracer, zipkinTracer)

	g := MountListeners(zapLogger, httpHandler)
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	zapLogger.Info("exit", zap.Any("exiting process", g.Run()))
}

func MountListeners(zapLogger *zap.Logger, httpHandler http.Handler) group.Group {
	// Now we're to the part of the func main where we want to start actually
	// running things, like servers bound to listeners to receive connections.
	//
	// The method is the same for each component: add a new actor to the group
	// struct, which is a combination of 2 anonymous functions: the first
	// function actually runs the component, and the second function should
	// interrupt the first function and cause it to return. It's in these
	// functions that we actually bind the Go kit server/handler structs to the
	// concrete transports and run them.
	//
	// Putting each component into its own block is mostly for aesthetics: it
	// clearly demarcates the scope in which each listener/socket may be used.
	var g group.Group
	{
		// The debug listener mounts the http.DefaultServeMux, and serves up
		// stuff like the Prometheus metrics route, the Go debug and profiling
		// routes, and so on.
		debugListener, err := net.Listen("tcp", config.Config.Debug)
		if err != nil {
			zapLogger.Info("service connected", zap.String("transport", "debug/HTTP"), zap.String("during", "Listen"), zap.Any("error", err))
			os.Exit(1)
		}
		g.Add(func() error {
			zapLogger.Info("service connected", zap.String("transport", "debug/HTTP"), zap.String("addr", config.Config.Debug))
			return http.Serve(debugListener, http.DefaultServeMux)
		}, func(error) {
			_ = debugListener.Close()
		})
	}
	{
		httpListener, err := net.Listen("tcp", config.Config.Http)
		if err != nil {
			zapLogger.Info("transport", zap.String("transport", "debug/HTTP"), zap.Any("err", err))
			os.Exit(1)
		}
		g.Add(func() error {
			zapLogger.Info("transport", zap.String("transport", "/HTTP"), zap.String("addr", config.Config.Http))
			return http.Serve(httpListener, httpHandler)
		}, func(error) {
			_ = httpListener.Close()
		})
	}
	return g
}

func InitOpenTracer(zipkinTracer *zipkin.Tracer, zapLogger *zap.Logger) (stdopentracing.Tracer, *zipkin.Tracer) {
	var tracer stdopentracing.Tracer
	{
		if config.Config.ZipkinBridge && zipkinTracer != nil {
			zapLogger.Info("Tracer", zap.String("type of tracer", "zipkin"),
				zap.String("URL", config.Config.ZipkinUrl))
			tracer = zipkinot.Wrap(zipkinTracer)
			zipkinTracer = nil // do not instrument with both native tracer and opentracing bridge
		} else if config.Config.LightstepToken != "" {
			zapLogger.Info("Tracer", zap.String("type of tracer", "LightStep"))
			tracer = lightstep.NewTracer(lightstep.Options{
				AccessToken: config.Config.LightstepToken,
			})
			defer lightstep.FlushLightStepTracer(tracer)
		} else if config.Config.Appdash != "" {
			zapLogger.Info("Tracer", zap.String("type of tracer", "Appdash"),
				zap.String("Appdash", config.Config.Appdash))
			tracer = appdashot.NewTracer(appdash.NewRemoteCollector(config.Config.Appdash))
		} else {
			tracer = stdopentracing.GlobalTracer() // no-op
		}
	}
	return tracer, zipkinTracer
}

func InitZipkinTracer(zapLogger *zap.Logger) *zipkin.Tracer {
	var zipkinTracer *zipkin.Tracer
	{
		if config.Config.ZipkinUrl != "" {
			var (
				err         error
				hostPort    = "8080"
				serviceName = config.Config.Name
				reporter    = zipkinhttp.NewReporter(config.Config.ZipkinUrl)
			)
			defer reporter.Close()
			zEP, err := zipkin.NewEndpoint(serviceName, hostPort)
			if err != nil {
				zapLogger.Error(err.Error())
				os.Exit(1)
			}

			sampler, err := zipkin.NewCountingSampler(1)
			if err != nil {
				zapLogger.Error(err.Error())
				os.Exit(1)
			}

			zipkinTracer, err = zipkin.NewTracer(reporter, zipkin.WithSampler(sampler), zipkin.WithLocalEndpoint(zEP))
			if err != nil {
				zapLogger.Error(err.Error())
				os.Exit(1)
			}
			if !(config.Config.ZipkinBridge) {
				zapLogger.Info("Tracer", zap.String("type of tracer", "zipkin"),
					zap.String("URL", config.Config.ZipkinUrl))
			}
		}
	}
	return zipkinTracer
}

// InitService builds the layers of the service "onion" from the inside out. First, the
// business logic service; then, the set of endpoints that wrap the service;
// and finally, a series of concrete transport adapters. The adapters, like
// the HTTP handler or the gRPC server, are the bridge between Go kit and
// the interfaces that the transports expect
func InitService(zapLogger *zap.Logger, db *gorm.DB, amqpproducerconn service.Queue,
	amqpconsumerconn service.Queue, counter service.Counters,
	tracer stdopentracing.Tracer, zipkinTracer *zipkin.Tracer) http.Handler {

	var (
		svc         = service.New(zapLogger, db, amqpproducerconn, amqpconsumerconn, counter)
		endpoints   = endpoint.New(svc, zapLogger, counter.Duration, tracer, zipkinTracer)
		httpHandler = transport.NewHTTPHandler(svc, endpoints, counter.Duration, tracer, zipkinTracer, zapLogger)
	)
	return httpHandler
}

// InitQueues initializes a set of producer and consumer amqp queues to be used for things such as
// account registration emails amongst many others.
func InitQueues(zapLogger *zap.Logger) (service.Queue, service.Queue) {
	amqpConnString := "amqp://user:bitnami@stats/"
	producerQueueNames := []string{"lens_welcome_email", "lens_password_reset_email", "lens_email_reset_email"}
	consumerQueueNames := []string{"user_inactive"}
	amqpproducerconn, err := service.NewAmqpConnection(amqpConnString, producerQueueNames)
	if err != nil {
		zapLogger.Error(err.Error())
	}
	amqpconsumerconn, err := service.NewAmqpConnection(amqpConnString, consumerQueueNames)
	if err != nil {
		zapLogger.Error(err.Error())
	}
	return amqpproducerconn, amqpconsumerconn
}

// InitMetrics Creates the (sparse) metrics used in the service.
func InitMetrics() service.Counters {
	var createUserReq, successfulCreateUserReq, failedCreateUserReq, getUserRequests, successfulGetUserReq,
		failedGetUserReq, successfulLogInReq, failedLogInReq metrics.Counter
	{
		// Business-level metrics.
		createUserReq = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "create_user_requests",
			Help:      "Total count of create user requests via the CreateUser method.",
		}, []string{})
		successfulCreateUserReq = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "create_user_success_ops",
			Help:      "Total count of successful create user requests via the CreateUser method.",
		}, []string{})
		failedCreateUserReq = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "create_user_failed_ops",
			Help:      "Total count of failed create user requests via the CreateUser method.",
		}, []string{})
		getUserRequests = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "get_user_requests",
			Help:      "Total count of get user requests.",
		}, []string{})
		successfulGetUserReq = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "get_user_requests_success_ops",
			Help:      "Total count of successful get user requests.",
		}, []string{})
		failedGetUserReq = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "get_user_requests_failed_ops",
			Help:      "Total count of failed get user requests.",
		}, []string{})
		successfulLogInReq = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "login_requests_sucess_ops",
			Help:      "Total count of successful logIn requests.",
		}, []string{})
		failedLogInReq = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "login_requests_failed_ops",
			Help:      "Total count of failed login requests.",
		}, []string{})
	}

	var duration metrics.Histogram
	{
		// Endpoint-level metrics.
		duration = prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "request_duration_seconds",
			Help:      "Request duration in seconds.",
		}, []string{"method", "success"})
	}

	counter := service.Counters{
		CreateUserRequest:           createUserReq,
		SuccessfulCreateUserRequest: successfulCreateUserReq,
		FailedCreateUserRequest:     failedCreateUserReq,
		GetUserRequest:              getUserRequests,
		SuccessfulGetUserRequest:    successfulGetUserReq,
		FailedGetUserRequest:        failedGetUserReq,
		SuccessfulLogInRequest:      successfulLogInReq,
		FailedLogInRequest:          failedLogInReq,
		Duration:                    duration,
	}

	return counter
}

// InitDbConnection initializes a database connection and creates associated tables/migrates schemas
func InitDbConnection(zapLogger *zap.Logger) (*gorm.DB, error) {
	connString := config.Config.GetDatabaseConnectionString()
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		zapLogger.Error(err.Error())
		os.Exit(1)
	}

	zapLogger.Info("successfully connected to database")
	db.SingularTable(true)
	db.LogMode(false)
	CreateTablesOrMigrateSchemas(db, zapLogger)

	return db, err
}

// CreateTablesOrMigrateSchemas creates a given set of tables based on a schema
// if it does not exist or migrates the table schemas to the latest version
func CreateTablesOrMigrateSchemas(db *gorm.DB, zapLogger *zap.Logger) {
	var userTable models.UserORM
	var teamsTable models.TeamORM
	var groupTable models.GroupORM
	userTable.MigrateSchemaOrCreateTable(db,zapLogger)
	teamsTable.MigrateSchemaOrCreateTable(db,zapLogger)
	groupTable.MigrateSchemaOrCreateTable(db,zapLogger)
}

// usageFor is used to parse Operating System Flags defined
func usageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		_, _ = fmt.Fprintf(os.Stderr, "USAGE\n")
		_, _ = fmt.Fprintf(os.Stderr, "  %s\n", short)
		_, _ = fmt.Fprintf(os.Stderr, "\n")
		_, _ = fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			_, _ = fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		_ = w.Flush()
		_, _ = fmt.Fprintf(os.Stderr, "\n")
	}
}
