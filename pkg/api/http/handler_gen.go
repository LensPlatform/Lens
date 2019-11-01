package http

import (
	http1 "net/http"

	"github.com/go-kit/kit/transport/http"

	"LensPlatform/Lens/pkg/api/endpoint"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := http1.NewServeMux()
	makeCreateUserAccountHandler(m, endpoints, options["CreateUserAccount"])
	makeUpdateUserAccountHandler(m, endpoints, options["UpdateUserAccount"])
	makeDeleteUserAccountHandler(m, endpoints, options["DeleteUserAccount"])
	makeGetUserAccountHandler(m, endpoints, options["GetUserAccount"])
	makeCreateTeamsAccountHandler(m, endpoints, options["CreateTeamsAccount"])
	makeUpdateTeamsAccountHandler(m, endpoints, options["UpdateTeamsAccount"])
	makeDeleteTeamsAccountHandler(m, endpoints, options["DeleteTeamsAccount"])
	makeGetTeamsAccountHandler(m, endpoints, options["GetTeamsAccount"])
	makeAddUserToTeamsAccountHandler(m, endpoints, options["AddUserToTeamsAccount"])
	makeDeleteUserFromTeamsAccountHandler(m, endpoints, options["DeleteUserFromTeamsAccount"])
	makeCreateGroupHandler(m, endpoints, options["CreateGroup"])
	makeSubscribeToGroupHandler(m, endpoints, options["SubscribeToGroup"])
	makeUnsubscribeFromGroupHandler(m, endpoints, options["UnsubscribeFromGroup"])
	makeDeleteGroupHandler(m, endpoints, options["DeleteGroup"])
	makeUpdateGroupHandler(m, endpoints, options["UpdateGroup"])
	makeGetGroupByIDHandler(m, endpoints, options["GetGroupByID"])
	makeGetGroupByNameHandler(m, endpoints, options["GetGroupByName"])
	makeIsGroupPrivateHandler(m, endpoints, options["IsGroupPrivate"])
	return m
}