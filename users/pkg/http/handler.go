package http

import (
	"context"
	"encoding/json"
	"errors"
	endpoint "github.com/LensPlatform/Lens/users/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	"net/http"
)

// makeCreateUserAccountHandler creates the handler logic
func makeCreateUserAccountHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-user-account", http1.NewServer(endpoints.CreateUserAccountEndpoint, decodeCreateUserAccountRequest, encodeCreateUserAccountResponse, options...))
}

// decodeCreateUserAccountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateUserAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateUserAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateUserAccountResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateUserAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateUserAccountHandler creates the handler logic
func makeUpdateUserAccountHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-user-account", http1.NewServer(endpoints.UpdateUserAccountEndpoint, decodeUpdateUserAccountRequest, encodeUpdateUserAccountResponse, options...))
}

// decodeUpdateUserAccountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateUserAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateUserAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateUserAccountResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateUserAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteUserAccountHandler creates the handler logic
func makeDeleteUserAccountHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-user-account", http1.NewServer(endpoints.DeleteUserAccountEndpoint, decodeDeleteUserAccountRequest, encodeDeleteUserAccountResponse, options...))
}

// decodeDeleteUserAccountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteUserAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteUserAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteUserAccountResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteUserAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetUserAccountHandler creates the handler logic
func makeGetUserAccountHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-user-account", http1.NewServer(endpoints.GetUserAccountEndpoint, decodeGetUserAccountRequest, encodeGetUserAccountResponse, options...))
}

// decodeGetUserAccountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUserAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetUserAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetUserAccountResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUserAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateTeamsAccountHandler creates the handler logic
func makeCreateTeamsAccountHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-teams-account", http1.NewServer(endpoints.CreateTeamsAccountEndpoint, decodeCreateTeamsAccountRequest, encodeCreateTeamsAccountResponse, options...))
}

// decodeCreateTeamsAccountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateTeamsAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateTeamsAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateTeamsAccountResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateTeamsAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateTeamsAccountHandler creates the handler logic
func makeUpdateTeamsAccountHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-teams-account", http1.NewServer(endpoints.UpdateTeamsAccountEndpoint, decodeUpdateTeamsAccountRequest, encodeUpdateTeamsAccountResponse, options...))
}

// decodeUpdateTeamsAccountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateTeamsAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateTeamsAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateTeamsAccountResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateTeamsAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteTeamsAccountHandler creates the handler logic
func makeDeleteTeamsAccountHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-teams-account", http1.NewServer(endpoints.DeleteTeamsAccountEndpoint, decodeDeleteTeamsAccountRequest, encodeDeleteTeamsAccountResponse, options...))
}

// decodeDeleteTeamsAccountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteTeamsAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteTeamsAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteTeamsAccountResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteTeamsAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetTeamsAccountHandler creates the handler logic
func makeGetTeamsAccountHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-teams-account", http1.NewServer(endpoints.GetTeamsAccountEndpoint, decodeGetTeamsAccountRequest, encodeGetTeamsAccountResponse, options...))
}

// decodeGetTeamsAccountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetTeamsAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetTeamsAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetTeamsAccountResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetTeamsAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddUserToTeamsAccountHandler creates the handler logic
func makeAddUserToTeamsAccountHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/add-user-to-teams-account", http1.NewServer(endpoints.AddUserToTeamsAccountEndpoint, decodeAddUserToTeamsAccountRequest, encodeAddUserToTeamsAccountResponse, options...))
}

// decodeAddUserToTeamsAccountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddUserToTeamsAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.AddUserToTeamsAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAddUserToTeamsAccountResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddUserToTeamsAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteUserFromTeamsAccountHandler creates the handler logic
func makeDeleteUserFromTeamsAccountHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-user-from-teams-account", http1.NewServer(endpoints.DeleteUserFromTeamsAccountEndpoint, decodeDeleteUserFromTeamsAccountRequest, encodeDeleteUserFromTeamsAccountResponse, options...))
}

// decodeDeleteUserFromTeamsAccountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteUserFromTeamsAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteUserFromTeamsAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteUserFromTeamsAccountResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteUserFromTeamsAccountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateGroupHandler creates the handler logic
func makeCreateGroupHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-group", http1.NewServer(endpoints.CreateGroupEndpoint, decodeCreateGroupRequest, encodeCreateGroupResponse, options...))
}

// decodeCreateGroupRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateGroupRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateGroupResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateGroupResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSubscribeToGroupHandler creates the handler logic
func makeSubscribeToGroupHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/subscribe-to-group", http1.NewServer(endpoints.SubscribeToGroupEndpoint, decodeSubscribeToGroupRequest, encodeSubscribeToGroupResponse, options...))
}

// decodeSubscribeToGroupRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSubscribeToGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SubscribeToGroupRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSubscribeToGroupResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSubscribeToGroupResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUnsubscribeFromGroupHandler creates the handler logic
func makeUnsubscribeFromGroupHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/unsubscribe-from-group", http1.NewServer(endpoints.UnsubscribeFromGroupEndpoint, decodeUnsubscribeFromGroupRequest, encodeUnsubscribeFromGroupResponse, options...))
}

// decodeUnsubscribeFromGroupRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUnsubscribeFromGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UnsubscribeFromGroupRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUnsubscribeFromGroupResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUnsubscribeFromGroupResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteGroupHandler creates the handler logic
func makeDeleteGroupHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-group", http1.NewServer(endpoints.DeleteGroupEndpoint, decodeDeleteGroupRequest, encodeDeleteGroupResponse, options...))
}

// decodeDeleteGroupRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteGroupRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteGroupResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteGroupResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateGroupHandler creates the handler logic
func makeUpdateGroupHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-group", http1.NewServer(endpoints.UpdateGroupEndpoint, decodeUpdateGroupRequest, encodeUpdateGroupResponse, options...))
}

// decodeUpdateGroupRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateGroupRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateGroupResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateGroupResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetGroupByIDHandler creates the handler logic
func makeGetGroupByIDHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-group-by-id", http1.NewServer(endpoints.GetGroupByIDEndpoint, decodeGetGroupByIDRequest, encodeGetGroupByIDResponse, options...))
}

// decodeGetGroupByIDRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetGroupByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetGroupByIDRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetGroupByIDResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetGroupByIDResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetGroupByNameHandler creates the handler logic
func makeGetGroupByNameHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-group-by-name", http1.NewServer(endpoints.GetGroupByNameEndpoint, decodeGetGroupByNameRequest, encodeGetGroupByNameResponse, options...))
}

// decodeGetGroupByNameRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetGroupByNameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetGroupByNameRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetGroupByNameResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetGroupByNameResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeIsGroupPrivateHandler creates the handler logic
func makeIsGroupPrivateHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/is-group-private", http1.NewServer(endpoints.IsGroupPrivateEndpoint, decodeIsGroupPrivateRequest, encodeIsGroupPrivateResponse, options...))
}

// decodeIsGroupPrivateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeIsGroupPrivateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.IsGroupPrivateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeIsGroupPrivateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeIsGroupPrivateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
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
