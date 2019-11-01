package servicestub

import (
	_ "net"

	"github.com/go-kit/kit/log"

	endpoint1 "github.com/go-kit/kit/endpoint"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	_ "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	prometheus1 "github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"

	endpoint "LensPlatform/Lens/pkg/api/endpoint"
	http1 "LensPlatform/Lens/pkg/api/http"
	"LensPlatform/Lens/pkg/api/service"
)

func DefaultHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"AddUserToTeamsAccount":      {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "AddUserToTeamsAccount", logger))},
		"CreateGroup":                {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateGroup", logger))},
		"CreateTeamsAccount":         {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateTeamsAccount", logger))},
		"CreateUserAccount":          {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateUserAccount", logger))},
		"DeleteGroup":                {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteGroup", logger))},
		"DeleteTeamsAccount":         {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteTeamsAccount", logger))},
		"DeleteUserAccount":          {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteUserAccount", logger))},
		"DeleteUserFromTeamsAccount": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteUserFromTeamsAccount", logger))},
		"GetGroupByID":               {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetGroupByID", logger))},
		"GetGroupByName":             {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetGroupByName", logger))},
		"GetTeamsAccount":            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetTeamsAccount", logger))},
		"GetUserAccount":             {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetUserAccount", logger))},
		"IsGroupPrivate":             {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "IsGroupPrivate", logger))},
		"SubscribeToGroup":           {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "SubscribeToGroup", logger))},
		"UnsubscribeFromGroup":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UnsubscribeFromGroup", logger))},
		"UpdateGroup":                {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateGroup", logger))},
		"UpdateTeamsAccount":         {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateTeamsAccount", logger))},
		"UpdateUserAccount":          {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateUserAccount", logger))},
	}
	return options
}

func addDefaultEndpointMiddleware(logger *zap.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["CreateUserAccount"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "CreateUserAccount"))}
	mw["UpdateUserAccount"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "UpdateUserAccount"))}
	mw["DeleteUserAccount"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "DeleteUserAccount"))}
	mw["GetUserAccount"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "GetUserAccount"))}
	mw["CreateTeamsAccount"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "CreateTeamsAccount"))}
	mw["UpdateTeamsAccount"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "UpdateTeamsAccount"))}
	mw["DeleteTeamsAccount"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "DeleteTeamsAccount"))}
	mw["GetTeamsAccount"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "GetTeamsAccount"))}
	mw["AddUserToTeamsAccount"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "AddUserToTeamsAccount"))}
	mw["DeleteUserFromTeamsAccount"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "DeleteUserFromTeamsAccount"))}
	mw["CreateGroup"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "CreateGroup"))}
	mw["SubscribeToGroup"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "SubscribeToGroup"))}
	mw["UnsubscribeFromGroup"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "UnsubscribeFromGroup"))}
	mw["DeleteGroup"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "DeleteGroup"))}
	mw["UpdateGroup"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "UpdateGroup"))}
	mw["GetGroupByID"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "GetGroupByID"))}
	mw["GetGroupByName"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "GetGroupByName"))}
	mw["IsGroupPrivate"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(logger), endpoint.InstrumentingMiddleware(duration.With("method", "IsGroupPrivate"))}
}

func addDefaultServiceMiddleware(logger *zap.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}

func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"CreateUserAccount", "UpdateUserAccount", "DeleteUserAccount", "GetUserAccount", "CreateTeamsAccount", "UpdateTeamsAccount", "DeleteTeamsAccount", "GetTeamsAccount", "AddUserToTeamsAccount", "DeleteUserFromTeamsAccount", "CreateGroup", "SubscribeToGroup", "UnsubscribeFromGroup", "DeleteGroup", "UpdateGroup", "GetGroupByID", "GetGroupByName", "IsGroupPrivate"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}

func GetServiceMiddleware(logger *zap.Logger) (mw []service.Middleware) {
	mw = []service.Middleware{}
	mw = addDefaultServiceMiddleware(logger, mw)
	return
}
func GetEndpointMiddleware(logger *zap.Logger) (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{}
	duration := prometheus.NewSummaryFrom(prometheus1.SummaryOpts{
		Help:      "Request duration in seconds.",
		Name:      "request_duration_seconds",
		Namespace: "example",
		Subsystem: "users",
	}, []string{"method", "success"})
	addDefaultEndpointMiddleware(logger, duration, mw)
	return
}