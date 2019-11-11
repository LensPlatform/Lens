package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/go-kit/kit/metrics"
	"github.com/gorilla/mux"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	_ "github.com/go-kit/kit/log"
	_ "github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/tracing/zipkin"
	_ "github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"

	_ "github.com/go-kit/kit/endpoint"

	serviceendpoint "github.com/LensPlatform/Lens/pkg/endpoint"
	service "github.com/LensPlatform/Lens/pkg/service"
	utils "github.com/LensPlatform/Lens/pkg/helper"

)

// NewHTTPHandler returns an HTTP handler that makes a set of endpoints
// available on predefined paths.
func NewHTTPHandler(s service.Service, endpoints serviceendpoint.Set,
					duration metrics.Histogram, otTracer stdopentracing.Tracer,
					zipkinTracer *stdzipkin.Tracer, logger *zap.Logger) http.Handler {
	r := mux.NewRouter()
	e := serviceendpoint.MakeServerEndpoints(s,logger, duration, otTracer, zipkinTracer)
	var options = []httptransport.ServerOption{
		httptransport.ServerErrorHandler(utils.NewTransportHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}
	if zipkinTracer != nil {
		// Zipkin HTTP Server Trace can either be instantiated per endpoint with a
		// provided operation name or a global tracing service can be instantiated
		// without an operation name and fed to each Go kit endpoint as ServerOption.
		// In the latter case, the operation name will be the endpoint's http method.
		// We demonstrate a global tracing service here.
		options = append(options, zipkin.HTTPServerTrace(zipkinTracer))
	}

	// POST    /user/create-user                          creates a user profile
	// GET    /user/username/:username                         Gets a user profile by username
	// GET    /user/:id                         Gets a user profile by id
	// GET    /user/email/:email                         Gets a user profile by email
	r.Methods("POST").Path("/v1/user/create-user").Handler(httptransport.NewServer(
		e.CreateUserEndpoint,
		decodeCreateUserRequest,
		encodeResponse,
		options...,
		))
	r.Methods("GET").Path("/v1/user/username/{username}").Handler(httptransport.NewServer(
		e.GetUserByUsernameEndpoint,
		decodeGetUserRequestByUsername,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/v1/user/id/{id}").Handler(httptransport.NewServer(
		e.GetUserByIdEndpoint,
		decodeGetUserRequestById,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/v1/user/email/{email}").Handler(httptransport.NewServer(
		e.GetUserByEmailEndpoint,
		decodeGetUserRequestByEmail,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/v1/user/login").Handler(httptransport.NewServer(
		e.LoginEndpoint,
		decodeLogin,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/v1/metrics").Handler(promhttp.Handler())
	return r
}

// errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error. For more information, read the
// big comment in endpoints.go.
type errorer interface {
	error() error
}

// encodeResponse is the common method to encode all response types to the
// client. I chose to do it this way because, since we're using JSON, there's no
// reason to provide anything more specific. It's certainly possible to
// specialize on a per-response (per-method) basis.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return err
	}
	return nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case utils.ErrNotFound:
		return http.StatusNotFound
	case utils.ErrAlreadyExists, utils.ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

// decodeHTTPSumRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded CreateUser request from the HTTP request body. Primarily useful in a
// server.
func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req serviceendpoint.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetUserRequestById(_ context.Context, r *http.Request) (interface{}, error) {
	req, err := decodeGetUserRequest(r, "id")
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetUserRequestByEmail(_ context.Context, r *http.Request) (interface{}, error) {
	req, err := decodeGetUserRequest(r, "email")
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetUserRequestByUsername(_ context.Context, r *http.Request) (interface{}, error) {
	req, err := decodeGetUserRequest(r, "username")
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeLogin(_ context.Context, r *http.Request) (interface{}, error) {
	var req serviceendpoint.LoginRequest
	params := r.URL.Query()
	req.Username = params.Get("username")
	req.Password = params.Get("password")
	return req, nil
}

func decodeGetUserRequest(r *http.Request, param string) (interface{}, error) {
	var req serviceendpoint.GetUserRequest
	vars := mux.Vars(r)

	value, ok := vars[param]
	if !ok {
		return nil, utils.ErrBadRouting
	}
	req.Param = value
	return req, nil
}