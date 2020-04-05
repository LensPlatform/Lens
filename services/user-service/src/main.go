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
package src

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
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/oklog/oklog/pkg/group"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/LensPlatform/Lens/services/user-service/src/pkg/config"
	"github.com/LensPlatform/Lens/services/user-service/src/pkg/database/postgresql"
	"github.com/LensPlatform/Lens/services/user-service/src/pkg/endpoint"
	"github.com/LensPlatform/Lens/services/user-service/src/pkg/log"
	"github.com/LensPlatform/Lens/services/user-service/src/pkg/monitoring"
	"github.com/LensPlatform/Lens/services/user-service/src/pkg/queues"
	"github.com/LensPlatform/Lens/services/user-service/src/pkg/service"
	"github.com/LensPlatform/Lens/services/user-service/src/pkg/transport"
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

	tracer, zipkinTracer, counters := monitoring.Init(zapLogger)

	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())

	// configure sql db connection
	db, err := postgresql.Init(zapLogger)
	if err != nil {
		zapLogger.Error(err.Error(), zap.String("Connection Error", "Unable To Connect To Database"))
		os.Exit(1)
	}
	defer db.Close()

	// connect to rabbitmq
	amqpproducerconn, amqpconsumerconn := queues.Init(zapLogger)

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

// InitService builds the layers of the service "onion" from the inside out. First, the
// business logic service; then, the set of endpoints that wrap the service;
// and finally, a series of concrete transport adapters. The adapters, like
// the HTTP handler or the gRPC server, are the bridge between Go kit and
// the interfaces that the transports expect
func InitService(zapLogger *zap.Logger, db *gorm.DB, amqpproducerconn queues.Queue,
	amqpconsumerconn queues.Queue, counter service.Counters,
	tracer stdopentracing.Tracer, zipkinTracer *zipkin.Tracer) http.Handler {

	var (
		svc         = service.New(zapLogger, db, amqpproducerconn, amqpconsumerconn, counter)
		endpoints   = endpoint.New(svc, zapLogger, counter.Duration, tracer, zipkinTracer)
		httpHandler = transport.NewHTTPHandler(svc, endpoints, counter.Duration, tracer, zipkinTracer, zapLogger)
	)
	return httpHandler
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
