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

	"github.com/jmoiron/sqlx"
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
	"go.uber.org/zap/zapcore"
	"sourcegraph.com/sourcegraph/appdash"
	appdashot "sourcegraph.com/sourcegraph/appdash/opentracing"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"

	"github.com/LensPlatform/Lens/pkg/transport"

	"github.com/LensPlatform/Lens/pkg/endpoint"
	"github.com/LensPlatform/Lens/pkg/service"
)


func main() {
	// Define our flags. Your service probably won't need to bind listeners for
	// *all* supported transports, or support both Zipkin and LightStep, and so
	// on, but we do it here for demonstration purposes.
	fs := flag.NewFlagSet("svc", flag.ExitOnError)
	var (
		debugAddr      = fs.String("debug.addr", ":8084", "Debug and metrics listen address")
		httpAddr       = fs.String("http-addr", ":8085", "HTTP listen address")
		zipkinURL      = fs.String("zipkin-url", "", "Enable Zipkin tracing via HTTP reporter URL e.g. http://localhost:9411/api/v2/spans")
		zipkinBridge   = fs.Bool("zipkin-ot-bridge", false, "Use Zipkin OpenTracing bridge instead of native implementation")
		lightstepToken = fs.String("lightstep-token", "", "Enable LightStep tracing via a LightStep access token")
		appdashAddr    = fs.String("appdash-addr", "", "Enable Appdash tracing via an Appdash server host:port")
	)
	fs.Usage = usageFor(fs, os.Args[0]+" [flags]")
	_ = fs.Parse(os.Args[1:])

	// configure logging
	zapLogger, _ := initZap(viper.GetString("level"))
	defer zapLogger.Sync()
	stdLog := zap.RedirectStdLog(zapLogger)
	defer stdLog()

	var zipkinTracer *zipkin.Tracer
	{
		if *zipkinURL != "" {
			var (
				err         error
				hostPort    = "localhost:80"
				serviceName = "user"
				reporter    = zipkinhttp.NewReporter(*zipkinURL)
			)
			defer reporter.Close()
			zEP, _ := zipkin.NewEndpoint(serviceName, hostPort)
			zipkinTracer, err = zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(zEP))
			if err != nil {
				zapLogger.Error(err.Error())
				os.Exit(1)
			}
			if !(*zipkinBridge) {
				zapLogger.Info("Tracer", zap.String("type of tracer", "zipkin"),
								zap.String("URL", *zipkinURL))
			}
		}
	}

	// Determine which OpenTracing tracer to use. We'll pass the tracer to all the
	// components that use it, as a dependency.
	var tracer stdopentracing.Tracer
	{
		if *zipkinBridge && zipkinTracer != nil {
			zapLogger.Info("Tracer", zap.String("type of tracer", "zipkin"),
				zap.String("URL", *zipkinURL))
			tracer = zipkinot.Wrap(zipkinTracer)
			zipkinTracer = nil // do not instrument with both native tracer and opentracing bridge
		} else if *lightstepToken != "" {
			zapLogger.Info("Tracer", zap.String("type of tracer", "LightStep"))
			tracer = lightstep.NewTracer(lightstep.Options{
				AccessToken: *lightstepToken,
			})
			defer lightstep.FlushLightStepTracer(tracer)
		} else if *appdashAddr != "" {
			zapLogger.Info("Tracer", zap.String("type of tracer", "Appdash"),
				zap.String("Appdash", *appdashAddr))
			tracer = appdashot.NewTracer(appdash.NewRemoteCollector(*appdashAddr))
		} else {
			tracer = stdopentracing.GlobalTracer() // no-op
		}
	}

	// Create the (sparse) metrics we'll use in the service. They, too, are
	// dependencies that we pass to components that use them.
	var  CreateUserRequest, successfulCreateUserReq, failedCreateUserReq, getUserRequests, successfulGetUserReq,
	failedGetUserReq, successfulLogInReq,failedLogInReq metrics.Counter
	{
		// Business-level metrics.
		CreateUserRequest = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
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
	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())

	// configure sql db connection
	connString := "postgresql://doadmin:x9nec6ffkm1i3187@backend-datastore-do-user-6612421-0.db.ondigitalocean.com:25060/defaultdb?sslmode=require"
	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		zapLogger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	// Check if DB connection can be made, only for logging purposes, should not fail/exit
	err = db.Ping()
	if err != nil {
		zapLogger.Error("error", zap.Any("unable to connect to database", err))
	}

	zapLogger.Info("successfully connected to database",)

	// connect to rabbitmq
	amqpConnString := "amqp://user:bitnami@stats/"
	producerQueueNames := []string{"lens_welcome_email", "lens_password_reset_email", "lens_email_reset_email"}
	consumerQueueNames := []string{"user_inactive"}
	amqpproducerconn, err:= service.NewAmqpConnection(amqpConnString, producerQueueNames)

	if err != nil {
		zapLogger.Error(err.Error())
	}
	amqpconsumerconn ,err:= service.NewAmqpConnection(amqpConnString, consumerQueueNames)

	if err != nil {
		zapLogger.Error(err.Error())
	}
	// Build the layers of the service "onion" from the inside out. First, the
	// business logic service; then, the set of endpoints that wrap the service;
	// and finally, a series of concrete transport adapters. The adapters, like
	// the HTTP handler or the gRPC server, are the bridge between Go kit and
	// the interfaces that the transports expect. Note that we're not binding
	// them to ports or anything yet; we'll do that next.
	var (
		userservice        = service.New(zapLogger, db, amqpproducerconn, amqpconsumerconn, CreateUserRequest, successfulCreateUserReq,
										 failedCreateUserReq, getUserRequests, successfulGetUserReq,
										 failedGetUserReq, successfulLogInReq,failedLogInReq)
		endpoints      = endpoint.New(userservice, zapLogger, duration, tracer, zipkinTracer)
		httpHandler    = transport.NewHTTPHandler(userservice,endpoints,duration, tracer, zipkinTracer, zapLogger)
	)

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
		debugListener, err := net.Listen("tcp", *debugAddr)
		if err != nil {
			zapLogger.Info("service connected",zap.String("transport", "debug/HTTP"), zap.String("during", "Listen"), zap.Any("error", err))
			os.Exit(1)
		}
		g.Add(func() error {
			zapLogger.Info("service connected",zap.String("transport", "debug/HTTP"), zap.String("addr", *debugAddr))
			return http.Serve(debugListener, http.DefaultServeMux)
		}, func(error) {
			debugListener.Close()
		})
	}
	{
		// The HTTP listener mounts the Go kit HTTP handler we created.
		httpListener, err := net.Listen("tcp", *httpAddr)
		if err != nil {
			zapLogger.Info("transport",zap.String("transport", "debug/HTTP"), zap.Any("err", err))
			os.Exit(1)
		}
		g.Add(func() error {
			zapLogger.Info("transport",zap.String("transport", "/HTTP"), zap.String("addr", *httpAddr))
			return http.Serve(httpListener, httpHandler)
		}, func(error) {
			_ = httpListener.Close()
		})
	}
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

func initZap(logLevel string) (*zap.Logger, error) {
	level := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	switch logLevel {
	case "debug":
		level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "info":
		level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "fatal":
		level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
	case "panic":
		level = zap.NewAtomicLevelAt(zapcore.PanicLevel)
	}

	zapEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	zapConfig := zap.Config{
		Level:       level,
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zapEncoderConfig,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	return zapConfig.Build()
}


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