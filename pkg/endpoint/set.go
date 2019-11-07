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


// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them
// into an Endpoints, and return it to the caller as a Service.
type Set struct {
	CreateUserEndpoint endpoint.Endpoint
	GetUserByIdEndpoint endpoint.Endpoint
	GetUserByUsernameEndpoint endpoint.Endpoint
	GetUserByEmailEndpoint endpoint.Endpoint
}

// New returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func New(svc service.Service, logger *zap.Logger, duration metrics.Histogram, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer) Set {
	return MakeServerEndpoints(svc, logger, duration, otTracer, zipkinTracer)
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service. Useful in a profilesvc
// server.
func MakeServerEndpoints(s service.Service, logger *zap.Logger,
	duration metrics.Histogram, otTracer stdopentracing.Tracer,
	zipkinTracer *stdzipkin.Tracer) Set {
	return Set{
		CreateUserEndpoint:   MakeCreateUserEndpoint(s, logger, duration, otTracer, zipkinTracer, "CreateUser"),
		GetUserByIdEndpoint:  MakeGetUserByIdEndpoint(s, logger, duration, otTracer, zipkinTracer, "GetUserById"),
		GetUserByUsernameEndpoint: MakeGetUserByUsernameEndpoint(s, logger, duration, otTracer, zipkinTracer, "GetUserByUsername"),
		GetUserByEmailEndpoint: MakeGetUserByEmailEndpoint(s, logger, duration, otTracer, zipkinTracer, "GetUserByEmail"),
	}
}

// ============================== Endpoint Definitions ======================
// CreateUserEndpoint constructs a Sum endpoint wrapping the service.
func MakeCreateUserEndpoint(s service.Service, logger *zap.Logger,
	duration metrics.Histogram, otTracer stdopentracing.Tracer,
	zipkinTracer *stdzipkin.Tracer, operationName string) endpoint.Endpoint {

		createUserEndpoint := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateUserRequest)
		logger.Info("User", zap.Any("user requesting creation", request))
		err = s.CreateUser(ctx, req.User)
		if err != nil {
			logger.Error(err.Error())
		}
		return CreateUserResponse{Err: err}, nil
	}
	return WrapMiddlewares(createUserEndpoint, logger,
			duration, otTracer, zipkinTracer, operationName)
}

func MakeGetUserByIdEndpoint(s service.Service, logger *zap.Logger,
	duration metrics.Histogram, otTracer stdopentracing.Tracer,
	zipkinTracer *stdzipkin.Tracer, operationName string) endpoint.Endpoint {

	getUserByIdEndpoint := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserRequest)
		logger.Info("User", zap.Any("attempting to get user by id", request))
		user, err := s.GetUserById(ctx, req.Param)
		if err != nil {
			logger.Error(err.Error())
		}
		return GetUserResponse{Err: err, User:user}, nil
	}
	return WrapMiddlewares(getUserByIdEndpoint, logger,
		duration, otTracer, zipkinTracer, operationName)
}

func MakeGetUserByUsernameEndpoint(s service.Service, logger *zap.Logger,
	duration metrics.Histogram, otTracer stdopentracing.Tracer,
	zipkinTracer *stdzipkin.Tracer, operationName string) endpoint.Endpoint {

	getUserByUsernameEndpoint := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserRequest)
		logger.Info("User", zap.Any("attempting to get user by username", request))
		user, err := s.GetUserByUsername(ctx, req.Param)
		if err != nil {
			logger.Error(err.Error())
		}
		return GetUserResponse{Err: err, User:user}, nil
	}
	return WrapMiddlewares(getUserByUsernameEndpoint, logger,
		duration, otTracer, zipkinTracer, operationName)
}

func MakeGetUserByEmailEndpoint(s service.Service, logger *zap.Logger,
	duration metrics.Histogram, otTracer stdopentracing.Tracer,
	zipkinTracer *stdzipkin.Tracer, operationName string) endpoint.Endpoint {

	getUserByEmailEndpoint := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserRequest)
		logger.Info("User", zap.Any("attempting to get user by email", request))
		user, err := s.GetUserByEmail(ctx, req.Param)
		if err != nil {
			logger.Error(err.Error())
		}
		return GetUserResponse{Err: err, User:user}, nil
	}
	return WrapMiddlewares(getUserByEmailEndpoint, logger,
		duration, otTracer, zipkinTracer, operationName)
}

// ============================== Endpoint Service Interface Impl  ======================
// CreateUser implements the service interface so that set may be used as a service.
func (s Set) CreateUser(ctx context.Context, user service.User)(err error){
	resp, err := s.CreateUserEndpoint(ctx, CreateUserRequest{User:user})
	if err != nil {
		return err
	}
	response := resp.(CreateUserResponse)
	return response.Err
}

func (s Set) GetUserById(ctx context.Context, id string)(user service.User, err error){
	resp, err := s.GetUserByIdEndpoint(ctx, GetUserRequest{Param:id})
	response := resp.(GetUserResponse)
	if err != nil {
		return response.User, err
	}
	return response.User, nil
}

func (s Set) GetUserByEmail(ctx context.Context, email string)(user service.User, err error){
	resp, err := s.GetUserByEmailEndpoint(ctx, GetUserRequest{Param:email})
	response := resp.(GetUserResponse)
	if err != nil {
		return response.User, err
	}
	return response.User, nil
}

func (s Set) GetUserByUsername(ctx context.Context, username string)(user service.User, err error){
	resp, err := s.GetUserByUsernameEndpoint(ctx, GetUserRequest{Param:username})
	response := resp.(GetUserResponse)
	if err != nil {
		return response.User, err
	}
	return response.User, nil
}

func WrapMiddlewares(endpoint endpoint.Endpoint, logger *zap.Logger,
	duration metrics.Histogram, otTracer stdopentracing.Tracer,
	zipkinTracer *stdzipkin.Tracer, operationName string) endpoint.Endpoint{

	endpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(endpoint)
	endpoint = opentracing.TraceServer(otTracer, operationName)(endpoint)
	if zipkinTracer != nil {
		endpoint = zipkin.TraceEndpoint(zipkinTracer, operationName)(endpoint)
	}
	endpoint = LoggingMiddleware(logger)(endpoint)
	endpoint = InstrumentingMiddleware(duration.With("method", operationName))(endpoint)
	return endpoint
}

// ============================== Endpoint Fail Time Assertions ======================

// compile time assertions for our response types implementing endpoint.Failer.
var (
	_ endpoint.Failer = CreateUserResponse{}
	_ endpoint.Failer = GetUserResponse{}
)

// ============================== Endpoint Request Definitions ======================

// CreateUserRequest collects the request parameters for the CreateUser method.
type CreateUserRequest struct {
	User service.User
}

type GetUserRequest struct {
	Param string
}

// ============================== Endpoint Response Definitions ======================

// CreateUserResponse collects the response values for the CreateUser method.
type CreateUserResponse struct {
	Err error `json:"err"` // should be intercepted by Failed/errorEncoder
}

type GetUserResponse struct {
	Err error `json:"err"`
	User service.User `json:"user"`
}

// ============================== Endpoint Response Failed Definitions ======================
func (r CreateUserResponse) error() error { return r.Err }
func (r CreateUserResponse) Failed() error { return r.Err }
func (r GetUserResponse) error() error { return r.Err }
func (r GetUserResponse) Failed() error { return r.Err }
