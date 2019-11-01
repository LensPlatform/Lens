package endpoint

import (
	service "github.com/LensPlatform/Lens/pkg/api/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateUserAccountEndpoint          endpoint.Endpoint
	UpdateUserAccountEndpoint          endpoint.Endpoint
	DeleteUserAccountEndpoint          endpoint.Endpoint
	GetUserAccountEndpoint             endpoint.Endpoint
	CreateTeamsAccountEndpoint         endpoint.Endpoint
	UpdateTeamsAccountEndpoint         endpoint.Endpoint
	DeleteTeamsAccountEndpoint         endpoint.Endpoint
	GetTeamsAccountEndpoint            endpoint.Endpoint
	AddUserToTeamsAccountEndpoint      endpoint.Endpoint
	DeleteUserFromTeamsAccountEndpoint endpoint.Endpoint
	CreateGroupEndpoint                endpoint.Endpoint
	SubscribeToGroupEndpoint           endpoint.Endpoint
	UnsubscribeFromGroupEndpoint       endpoint.Endpoint
	DeleteGroupEndpoint                endpoint.Endpoint
	UpdateGroupEndpoint                endpoint.Endpoint
	GetGroupByIDEndpoint               endpoint.Endpoint
	GetGroupByNameEndpoint             endpoint.Endpoint
	IsGroupPrivateEndpoint             endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.UsersService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		AddUserToTeamsAccountEndpoint:      MakeAddUserToTeamsAccountEndpoint(s),
		CreateGroupEndpoint:                MakeCreateGroupEndpoint(s),
		CreateTeamsAccountEndpoint:         MakeCreateTeamsAccountEndpoint(s),
		CreateUserAccountEndpoint:          MakeCreateUserAccountEndpoint(s),
		DeleteGroupEndpoint:                MakeDeleteGroupEndpoint(s),
		DeleteTeamsAccountEndpoint:         MakeDeleteTeamsAccountEndpoint(s),
		DeleteUserAccountEndpoint:          MakeDeleteUserAccountEndpoint(s),
		DeleteUserFromTeamsAccountEndpoint: MakeDeleteUserFromTeamsAccountEndpoint(s),
		GetGroupByIDEndpoint:               MakeGetGroupByIDEndpoint(s),
		GetGroupByNameEndpoint:             MakeGetGroupByNameEndpoint(s),
		GetTeamsAccountEndpoint:            MakeGetTeamsAccountEndpoint(s),
		GetUserAccountEndpoint:             MakeGetUserAccountEndpoint(s),
		IsGroupPrivateEndpoint:             MakeIsGroupPrivateEndpoint(s),
		SubscribeToGroupEndpoint:           MakeSubscribeToGroupEndpoint(s),
		UnsubscribeFromGroupEndpoint:       MakeUnsubscribeFromGroupEndpoint(s),
		UpdateGroupEndpoint:                MakeUpdateGroupEndpoint(s),
		UpdateTeamsAccountEndpoint:         MakeUpdateTeamsAccountEndpoint(s),
		UpdateUserAccountEndpoint:          MakeUpdateUserAccountEndpoint(s),
	}
	// Wrap each endpoint/function invocation in a midddleware
	for _, m := range mdw["CreateUserAccount"] {
		eps.CreateUserAccountEndpoint = m(eps.CreateUserAccountEndpoint)
	}
	for _, m := range mdw["UpdateUserAccount"] {
		eps.UpdateUserAccountEndpoint = m(eps.UpdateUserAccountEndpoint)
	}
	for _, m := range mdw["DeleteUserAccount"] {
		eps.DeleteUserAccountEndpoint = m(eps.DeleteUserAccountEndpoint)
	}
	for _, m := range mdw["GetUserAccount"] {
		eps.GetUserAccountEndpoint = m(eps.GetUserAccountEndpoint)
	}
	for _, m := range mdw["CreateTeamsAccount"] {
		eps.CreateTeamsAccountEndpoint = m(eps.CreateTeamsAccountEndpoint)
	}
	for _, m := range mdw["UpdateTeamsAccount"] {
		eps.UpdateTeamsAccountEndpoint = m(eps.UpdateTeamsAccountEndpoint)
	}
	for _, m := range mdw["DeleteTeamsAccount"] {
		eps.DeleteTeamsAccountEndpoint = m(eps.DeleteTeamsAccountEndpoint)
	}
	for _, m := range mdw["GetTeamsAccount"] {
		eps.GetTeamsAccountEndpoint = m(eps.GetTeamsAccountEndpoint)
	}
	for _, m := range mdw["AddUserToTeamsAccount"] {
		eps.AddUserToTeamsAccountEndpoint = m(eps.AddUserToTeamsAccountEndpoint)
	}
	for _, m := range mdw["DeleteUserFromTeamsAccount"] {
		eps.DeleteUserFromTeamsAccountEndpoint = m(eps.DeleteUserFromTeamsAccountEndpoint)
	}
	for _, m := range mdw["CreateGroup"] {
		eps.CreateGroupEndpoint = m(eps.CreateGroupEndpoint)
	}
	for _, m := range mdw["SubscribeToGroup"] {
		eps.SubscribeToGroupEndpoint = m(eps.SubscribeToGroupEndpoint)
	}
	for _, m := range mdw["UnsubscribeFromGroup"] {
		eps.UnsubscribeFromGroupEndpoint = m(eps.UnsubscribeFromGroupEndpoint)
	}
	for _, m := range mdw["DeleteGroup"] {
		eps.DeleteGroupEndpoint = m(eps.DeleteGroupEndpoint)
	}
	for _, m := range mdw["UpdateGroup"] {
		eps.UpdateGroupEndpoint = m(eps.UpdateGroupEndpoint)
	}
	for _, m := range mdw["GetGroupByID"] {
		eps.GetGroupByIDEndpoint = m(eps.GetGroupByIDEndpoint)
	}
	for _, m := range mdw["GetGroupByName"] {
		eps.GetGroupByNameEndpoint = m(eps.GetGroupByNameEndpoint)
	}
	for _, m := range mdw["IsGroupPrivate"] {
		eps.IsGroupPrivateEndpoint = m(eps.IsGroupPrivateEndpoint)
	}
	return eps
}
