package endpoint

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics"
	"github.com/sony/gobreaker"
	"go.uber.org/zap"
	"golang.org/x/time/rate"

	"github.com/LensPlatform/Lens/pkg/service"

	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"

	"github.com/go-kit/kit/circuitbreaker"

	_ "github.com/go-kit/kit/log"

	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/tracing/zipkin"
)

// Set collects all of the endpoints that compose an user service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Set struct {
	CreateUserEndpoint endpoint.Endpoint
}

// New returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func New(svc service.Service, logger *zap.Logger, duration metrics.Histogram, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer) Set {
	var createUserEndpoint endpoint.Endpoint
	{
		createUserEndpoint = MakeSumEndpoint(svc)
		createUserEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(createUserEndpoint)
		createUserEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(createUserEndpoint)
		createUserEndpoint = opentracing.TraceServer(otTracer, "CreateUser")(createUserEndpoint)
		if zipkinTracer != nil {
			createUserEndpoint = zipkin.TraceEndpoint(zipkinTracer, "CreateUser")(createUserEndpoint)
		}
		createUserEndpoint = LoggingMiddleware(logger)(createUserEndpoint)
		createUserEndpoint = InstrumentingMiddleware(duration.With("method", "CreateUser"))(createUserEndpoint)
	}

	return Set{
		CreateUserEndpoint:    createUserEndpoint,
	}
}

// ============================== Endpoint Service Interface Impl  ======================
// CreateUser implements the service interface so that set may be used as a service.
func (s Set) CreateUser(ctx context.Context, user interface{})(id string, err error){
	resp, err := s.CreateUserEndpoint(ctx, CreateUserRequest{user:user})
	if err != nil {
		return "", err
	}
	response := resp.(CreateUserResponse)
	return response.Id, response.Err
}

// ============================== Endpoint Definitions ======================
// MakeSumEndpoint constructs a Sum endpoint wrapping the service.
func MakeSumEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateUserRequest)
		id, err := s.CreateUser(ctx, req)
		return CreateUserResponse{Id: id, Err: err}, nil
	}
}

// ============================== Endpoint Fail Time Assertions ======================

// compile time assertions for our response types implementing endpoint.Failer.
var (
	_ endpoint.Failer = CreateUserResponse{}
)

// ============================== Endpoint Request Definitions ======================

// CreateUserRequest collects the request parameters for the CreateUser method.
type CreateUserRequest struct {
	user interface{}
}

// ============================== Endpoint Response Definitions ======================

// CreateUserResponse collects the response values for the CreateUser method.
type CreateUserResponse struct {
	Id   string   `json:"id"`
	Err error `json:"-"` // should be intercepted by Failed/errorEncoder
}

// ============================== Endpoint Response Failed Definitions ======================

// Failed implements endpoint.Failer
func (r CreateUserResponse) Failed() error { return r.Err }
