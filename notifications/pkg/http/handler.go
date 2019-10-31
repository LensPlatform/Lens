package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	endpoint "LensPlatform/Lens/notifications/pkg/endpoint"

	http1 "github.com/go-kit/kit/transport/http"
)

// makeSendEmailHandler creates the handler logic
func makeSendEmailHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/send-email", http1.NewServer(endpoints.SendEmailEndpoint, decodeSendEmailRequest, encodeSendEmailResponse, options...))
}

// decodeSendEmailRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSendEmailRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SendEmailRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSendEmailResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSendEmailResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSendWelcomeToLensEmailHandler creates the handler logic
func makeSendWelcomeToLensEmailHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/send-welcome-to-lens-email", http1.NewServer(endpoints.SendWelcomeToLensEmailEndpoint, decodeSendWelcomeToLensEmailRequest, encodeSendWelcomeToLensEmailResponse, options...))
}

// decodeSendWelcomeToLensEmailRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSendWelcomeToLensEmailRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SendWelcomeToLensEmailRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSendWelcomeToLensEmailResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSendWelcomeToLensEmailResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSendPasswordChangeEmailHandler creates the handler logic
func makeSendPasswordChangeEmailHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/send-password-change-email", http1.NewServer(endpoints.SendPasswordChangeEmailEndpoint, decodeSendPasswordChangeEmailRequest, encodeSendPasswordChangeEmailResponse, options...))
}

// decodeSendPasswordChangeEmailRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSendPasswordChangeEmailRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SendPasswordChangeEmailRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSendPasswordChangeEmailResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSendPasswordChangeEmailResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSendEmailAccountResetEmailHandler creates the handler logic
func makeSendEmailAccountResetEmailHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/send-email-account-reset-email", http1.NewServer(endpoints.SendEmailAccountResetEmailEndpoint, decodeSendEmailAccountResetEmailRequest, encodeSendEmailAccountResetEmailResponse, options...))
}

// decodeSendEmailAccountResetEmailRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSendEmailAccountResetEmailRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SendEmailAccountResetEmailRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSendEmailAccountResetEmailResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSendEmailAccountResetEmailResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
