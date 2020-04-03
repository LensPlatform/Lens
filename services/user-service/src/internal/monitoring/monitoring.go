package monitoring

import (
	"os"

	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"

	"github.com/lightstep/lightstep-tracer-go"
	stdopentracing "github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	"go.uber.org/zap"
	"sourcegraph.com/sourcegraph/appdash"
	appdashot "sourcegraph.com/sourcegraph/appdash/opentracing"

	"github.com/LensPlatform/Lens/services/user-service/src/pkg/config"
	"github.com/LensPlatform/Lens/services/user-service/src/pkg/service"
)

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
				hostPort    = config.Config.Http
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

func Init(logger *zap.Logger) (stdopentracing.Tracer, *zipkin.Tracer, service.Counters) {
	counters := InitMetrics()
	defaultZipkinTracer := InitZipkinTracer(logger)
	opentracing, zipkin := InitOpenTracer(defaultZipkinTracer, logger)
	return opentracing,zipkin,counters
}