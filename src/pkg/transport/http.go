/*
 * @File: transport.http.go
 * @Description: Defines REST endpoints for the user service
 * @Author: Yoan Yomba (yoanyomba@lens-platform.net)
 */
package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/metrics"
	"github.com/gorilla/mux"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/swaggo/swag"
	"go.uber.org/zap"
	"net/http"

	_ "github.com/go-kit/kit/log"
	_ "github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/tracing/zipkin"
	_ "github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"

	_ "github.com/go-kit/kit/endpoint"

	serviceendpoint "github.com/LensPlatform/Lens/src/pkg/endpoint"
	utils "github.com/LensPlatform/Lens/src/pkg/helper"
	service "github.com/LensPlatform/Lens/src/pkg/service"
)

// NewHTTPHandler returns an HTTP handler that makes a set of endpoints
// available on predefined paths.
func NewHTTPHandler(s service.Service, endpoints serviceendpoint.Set,
	duration metrics.Histogram, otTracer stdopentracing.Tracer,
	zipkinTracer *stdzipkin.Tracer, logger *zap.Logger) http.Handler {
	r := mux.NewRouter()
	e := serviceendpoint.MakeServerEndpoints(s, logger, duration, otTracer, zipkinTracer)
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

	CreateUserEndpoint(r, e, options)
	GetUserByUsername(r, e, options)
	GetUserById(r, e, options)
	LogInUser(r, e, options)
	GetServiceMetrics(r)
	GetSwaggerDocumentation(r, logger)

	return r
}

func GetSwaggerDocumentation(r *mux.Router, logger *zap.Logger) {
	r.Methods("GET").Path("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/docs/swagger/doc.json"),
	))
	r.Methods("GET").Path("/swagger.json").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		doc, err := swag.ReadDoc()
		if err != nil {
			logger.Error("swagger error", zap.Error(err), zap.String("path", "/docs/swagger.json"))
		}
		w.Write([]byte(doc))
	})
}

// GetServiceMetricsgodoc
// @Summary Prometheus metrics
// @Description returns HTTP requests duration and Go runtime metrics
// @Tags Kubernetes
// @Accept json
// @Produce json
// @Router /v1/metrics [get]
// @Success 200 {string} string "OK"
func GetServiceMetrics(r *mux.Router) *mux.Route {
	return r.Methods("GET").Path("/v1/metrics").Handler(promhttp.Handler())
}

// LogIn User godoc
// @Summary Hits the login user api endpoint
// @Description Attempts to log a user into the system given the user exists in the backend data store
// @Tags HTTP API
// @Accept json
// @Produce json
// @Router /v1/user/login [get]
// @Success 200
func LogInUser(r *mux.Router, e serviceendpoint.Set, options []httptransport.ServerOption) {
	GetUserByEmail(r, e, options)
	r.Methods("GET").Path("/v1/user/login").Handler(httptransport.NewServer(
		e.LoginEndpoint,
		decodeLogin,
		encodeResponse,
		options...,
	))
}

// Get User by Email godoc
// @Summary Hits the get user by email api endpoint
// @Description Obtains a user in the backend datastore based on the provided email
// @Tags HTTP API
// @Accept json
// @Produce json
// @Router /v1/user/email/{email} [get]
// @Success 200
func GetUserByEmail(r *mux.Router, e serviceendpoint.Set, options []httptransport.ServerOption) *mux.Route {
	return r.Methods("GET").Path("/v1/user/email/{email}").Handler(httptransport.NewServer(
		e.GetUserByEmailEndpoint,
		decodeGetUserRequestByEmail,
		encodeResponse,
		options...,
	))
}

// Get User by ID godoc
// @Summary Hits the get user by id api endpoint
// @Description Obtains a user in the backend datastore based on the provided id
// @Tags HTTP API
// @Accept json
// @Produce json
// @Router /v1/user/id/{id} [get]
// @Success 200
func GetUserById(r *mux.Router, e serviceendpoint.Set, options []httptransport.ServerOption) *mux.Route {
	return r.Methods("GET").Path("/v1/user/id/{id}").Handler(httptransport.NewServer(
		e.GetUserByIdEndpoint,
		decodeGetUserRequestById,
		encodeResponse,
		options...,
	))
}

// Get User by Username godoc
// @Summary Hits the get user by username api endpoint
// @Description Obtains a user in the backend datastore based on the provided username
// @Tags HTTP API
// @Accept json
// @Produce json
// @Router /v1/user/username/{username} [get]
// @Success 200
func GetUserByUsername(r *mux.Router, e serviceendpoint.Set, options []httptransport.ServerOption) *mux.Route {
	return r.Methods("GET").Path("/v1/user/username/{username}").Handler(httptransport.NewServer(
		e.GetUserByUsernameEndpoint,
		decodeGetUserRequestByUsername,
		encodeResponse,
		options...,
	))
}

// Create User godoc
// @Summary Hits the create user api endpoint
// @Description Creates a user in the backend datastore
// @Tags HTTP API
// @Accept json
// @Produce json
// @Router /v1/user/create [post]
// @Success 200
func CreateUserEndpoint(r *mux.Router, e serviceendpoint.Set, options []httptransport.ServerOption) *mux.Route {
	return r.Methods("POST").Path("/v1/user/create").Handler(httptransport.NewServer(
		e.CreateUserEndpoint,
		decodeCreateUserRequest,
		encodeResponse,
		options...,
	))
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
