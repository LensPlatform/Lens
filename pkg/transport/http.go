package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/tracing/zipkin"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"

	endpoint "github.com/go-kit/kit/endpoint"

	serviceendpoint "github.com/LensPlatform/Lens/pkg/endpoint"
	"github.com/LensPlatform/Lens/pkg/service"
)

// NewHTTPHandler returns an HTTP handler that makes a set of endpoints
// available on predefined paths.
func NewHTTPHandler(endpoints serviceendpoint.Set, otTracer stdopentracing.Tracer,
					zipkinTracer *stdzipkin.Tracer, logger log.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	if zipkinTracer != nil {
		// Zipkin HTTP Server Trace can either be instantiated per endpoint with a
		// provided operation name or a global tracing service can be instantiated
		// without an operation name and fed to each Go kit endpoint as ServerOption.
		// In the latter case, the operation name will be the endpoint's http method.
		// We demonstrate a global tracing service here.
		options = append(options, zipkin.HTTPServerTrace(zipkinTracer))
	}

	m := http.NewServeMux()
	m.Handle("/create-user", httptransport.NewServer(
		endpoints.CreateUserEndpoint,
		decodeHTTPCreateUserRequest,
		encodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "CreateUser", logger)))...,
	))
	return m
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	_ = json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func err2code(err error) int {
	switch err {
	case service.NullUser:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func errorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

type errorWrapper struct {
	Error string `json:"error"`
}

// decodeHTTPSumRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded CreateUser request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req serviceendpoint.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPSumResponse is a transport/http.DecodeResponseFunc that decodes a
// JSON-encoded CreateUser response from the HTTP response body. If the response has a
// non-200 status code, we will interpret that as an error and attempt to decode
// the specific error message from the response body. Primarily useful in a
// client.
func decodeHTTPCreateUserResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp serviceendpoint.CreateUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// encodeHTTPGenericRequest is a transport/http.EncodeRequestFunc that
// JSON-encodes any request to the request body. Primarily useful in a client.
func encodeHTTPGenericRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// encodeHTTPGenericResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}