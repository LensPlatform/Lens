package endpoint

import (
	"context"

	service "github.com/LensPlatform/Lens/users/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateUserAccountRequest collects the request parameters for the CreateUserAccount method.
type CreateUserAccountRequest struct {
	User     interface{} `json:"user"`
	UserType string      `json:"user_type"`
}

// CreateUserAccountResponse collects the response parameters for the CreateUserAccount method.
type CreateUserAccountResponse struct {
	E0 error `json:"e0"`
}

// MakeCreateUserAccountEndpoint returns an endpoint that invokes CreateUserAccount on the service.
func MakeCreateUserAccountEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserAccountRequest)
		e0 := s.CreateUserAccount(ctx, req.User, req.UserType)
		return CreateUserAccountResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r CreateUserAccountResponse) Failed() error {
	return r.E0
}

// UpdateUserAccountRequest collects the request parameters for the UpdateUserAccount method.
type UpdateUserAccountRequest struct {
	User interface{} `json:"user"`
	Id   string      `json:"id"`
}

// UpdateUserAccountResponse collects the response parameters for the UpdateUserAccount method.
type UpdateUserAccountResponse struct {
	E0 error `json:"e0"`
}

// MakeUpdateUserAccountEndpoint returns an endpoint that invokes UpdateUserAccount on the service.
func MakeUpdateUserAccountEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserAccountRequest)
		e0 := s.UpdateUserAccount(ctx, req.User, req.Id)
		return UpdateUserAccountResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r UpdateUserAccountResponse) Failed() error {
	return r.E0
}

// DeleteUserAccountRequest collects the request parameters for the DeleteUserAccount method.
type DeleteUserAccountRequest struct {
	Id string `json:"id"`
}

// DeleteUserAccountResponse collects the response parameters for the DeleteUserAccount method.
type DeleteUserAccountResponse struct {
	E0 error `json:"e0"`
}

// MakeDeleteUserAccountEndpoint returns an endpoint that invokes DeleteUserAccount on the service.
func MakeDeleteUserAccountEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserAccountRequest)
		e0 := s.DeleteUserAccount(ctx, req.Id)
		return DeleteUserAccountResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r DeleteUserAccountResponse) Failed() error {
	return r.E0
}

// GetUserAccountRequest collects the request parameters for the GetUserAccount method.
type GetUserAccountRequest struct {
	Id string `json:"id"`
}

// GetUserAccountResponse collects the response parameters for the GetUserAccount method.
type GetUserAccountResponse struct {
	User interface{} `json:"user"`
	Err  error       `json:"err"`
}

// MakeGetUserAccountEndpoint returns an endpoint that invokes GetUserAccount on the service.
func MakeGetUserAccountEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserAccountRequest)
		user, err := s.GetUserAccount(ctx, req.Id)
		return GetUserAccountResponse{
			Err:  err,
			User: user,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserAccountResponse) Failed() error {
	return r.Err
}

// CreateTeamsAccountRequest collects the request parameters for the CreateTeamsAccount method.
type CreateTeamsAccountRequest struct {
	Team     interface{} `json:"team"`
	TeamType string      `json:"team_type"`
}

// CreateTeamsAccountResponse collects the response parameters for the CreateTeamsAccount method.
type CreateTeamsAccountResponse struct {
	E0 error `json:"e0"`
}

// MakeCreateTeamsAccountEndpoint returns an endpoint that invokes CreateTeamsAccount on the service.
func MakeCreateTeamsAccountEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTeamsAccountRequest)
		e0 := s.CreateTeamsAccount(ctx, req.Team, req.TeamType)
		return CreateTeamsAccountResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r CreateTeamsAccountResponse) Failed() error {
	return r.E0
}

// UpdateTeamsAccountRequest collects the request parameters for the UpdateTeamsAccount method.
type UpdateTeamsAccountRequest struct {
	Team   interface{} `json:"team"`
	TeamID string      `json:"team_id"`
}

// UpdateTeamsAccountResponse collects the response parameters for the UpdateTeamsAccount method.
type UpdateTeamsAccountResponse struct {
	E0 error `json:"e0"`
}

// MakeUpdateTeamsAccountEndpoint returns an endpoint that invokes UpdateTeamsAccount on the service.
func MakeUpdateTeamsAccountEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateTeamsAccountRequest)
		e0 := s.UpdateTeamsAccount(ctx, req.Team, req.TeamID)
		return UpdateTeamsAccountResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r UpdateTeamsAccountResponse) Failed() error {
	return r.E0
}

// DeleteTeamsAccountRequest collects the request parameters for the DeleteTeamsAccount method.
type DeleteTeamsAccountRequest struct {
	TeamID string `json:"team_id"`
}

// DeleteTeamsAccountResponse collects the response parameters for the DeleteTeamsAccount method.
type DeleteTeamsAccountResponse struct {
	E0 error `json:"e0"`
}

// MakeDeleteTeamsAccountEndpoint returns an endpoint that invokes DeleteTeamsAccount on the service.
func MakeDeleteTeamsAccountEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteTeamsAccountRequest)
		e0 := s.DeleteTeamsAccount(ctx, req.TeamID)
		return DeleteTeamsAccountResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r DeleteTeamsAccountResponse) Failed() error {
	return r.E0
}

// GetTeamsAccountRequest collects the request parameters for the GetTeamsAccount method.
type GetTeamsAccountRequest struct {
	TeamID string `json:"team_id"`
}

// GetTeamsAccountResponse collects the response parameters for the GetTeamsAccount method.
type GetTeamsAccountResponse struct {
	Team interface{} `json:"team"`
	Err  error       `json:"err"`
}

// MakeGetTeamsAccountEndpoint returns an endpoint that invokes GetTeamsAccount on the service.
func MakeGetTeamsAccountEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTeamsAccountRequest)
		team, err := s.GetTeamsAccount(ctx, req.TeamID)
		return GetTeamsAccountResponse{
			Err:  err,
			Team: team,
		}, nil
	}
}

// Failed implements Failer.
func (r GetTeamsAccountResponse) Failed() error {
	return r.Err
}

// AddUserToTeamsAccountRequest collects the request parameters for the AddUserToTeamsAccount method.
type AddUserToTeamsAccountRequest struct {
	TeamID string `json:"team_id"`
	UserID string `json:"user_id"`
}

// AddUserToTeamsAccountResponse collects the response parameters for the AddUserToTeamsAccount method.
type AddUserToTeamsAccountResponse struct {
	E0 error `json:"e0"`
}

// MakeAddUserToTeamsAccountEndpoint returns an endpoint that invokes AddUserToTeamsAccount on the service.
func MakeAddUserToTeamsAccountEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddUserToTeamsAccountRequest)
		e0 := s.AddUserToTeamsAccount(ctx, req.TeamID, req.UserID)
		return AddUserToTeamsAccountResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r AddUserToTeamsAccountResponse) Failed() error {
	return r.E0
}

// DeleteUserFromTeamsAccountRequest collects the request parameters for the DeleteUserFromTeamsAccount method.
type DeleteUserFromTeamsAccountRequest struct {
	TeamID string `json:"team_id"`
	UserID string `json:"user_id"`
}

// DeleteUserFromTeamsAccountResponse collects the response parameters for the DeleteUserFromTeamsAccount method.
type DeleteUserFromTeamsAccountResponse struct {
	E0 error `json:"e0"`
}

// MakeDeleteUserFromTeamsAccountEndpoint returns an endpoint that invokes DeleteUserFromTeamsAccount on the service.
func MakeDeleteUserFromTeamsAccountEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserFromTeamsAccountRequest)
		e0 := s.DeleteUserFromTeamsAccount(ctx, req.TeamID, req.UserID)
		return DeleteUserFromTeamsAccountResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r DeleteUserFromTeamsAccountResponse) Failed() error {
	return r.E0
}

// CreateGroupRequest collects the request parameters for the CreateGroup method.
type CreateGroupRequest struct {
	GroupName string `json:"group_name"`
	GroupType string `json:"group_type"`
	IsPrivate bool   `json:"is_private"`
}

// CreateGroupResponse collects the response parameters for the CreateGroup method.
type CreateGroupResponse struct {
	E0 error `json:"e0"`
}

// MakeCreateGroupEndpoint returns an endpoint that invokes CreateGroup on the service.
func MakeCreateGroupEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateGroupRequest)
		e0 := s.CreateGroup(ctx, req.GroupName, req.GroupType, req.IsPrivate)
		return CreateGroupResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r CreateGroupResponse) Failed() error {
	return r.E0
}

// SubscribeToGroupRequest collects the request parameters for the SubscribeToGroup method.
type SubscribeToGroupRequest struct {
	GroupID string `json:"group_id"`
	ID      string `json:"id"`
}

// SubscribeToGroupResponse collects the response parameters for the SubscribeToGroup method.
type SubscribeToGroupResponse struct {
	E0 error `json:"e0"`
}

// MakeSubscribeToGroupEndpoint returns an endpoint that invokes SubscribeToGroup on the service.
func MakeSubscribeToGroupEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SubscribeToGroupRequest)
		e0 := s.SubscribeToGroup(ctx, req.GroupID, req.ID)
		return SubscribeToGroupResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r SubscribeToGroupResponse) Failed() error {
	return r.E0
}

// UnsubscribeFromGroupRequest collects the request parameters for the UnsubscribeFromGroup method.
type UnsubscribeFromGroupRequest struct {
	GroupID string `json:"group_id"`
	ID      string `json:"id"`
}

// UnsubscribeFromGroupResponse collects the response parameters for the UnsubscribeFromGroup method.
type UnsubscribeFromGroupResponse struct {
	E0 error `json:"e0"`
}

// MakeUnsubscribeFromGroupEndpoint returns an endpoint that invokes UnsubscribeFromGroup on the service.
func MakeUnsubscribeFromGroupEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UnsubscribeFromGroupRequest)
		e0 := s.UnsubscribeFromGroup(ctx, req.GroupID, req.ID)
		return UnsubscribeFromGroupResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r UnsubscribeFromGroupResponse) Failed() error {
	return r.E0
}

// DeleteGroupRequest collects the request parameters for the DeleteGroup method.
type DeleteGroupRequest struct {
	GroupID      string `json:"group_id"`
	GroupAdminID string `json:"group_admin_id"`
}

// DeleteGroupResponse collects the response parameters for the DeleteGroup method.
type DeleteGroupResponse struct {
	E0 error `json:"e0"`
}

// MakeDeleteGroupEndpoint returns an endpoint that invokes DeleteGroup on the service.
func MakeDeleteGroupEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteGroupRequest)
		e0 := s.DeleteGroup(ctx, req.GroupID, req.GroupAdminID)
		return DeleteGroupResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r DeleteGroupResponse) Failed() error {
	return r.E0
}

// UpdateGroupRequest collects the request parameters for the UpdateGroup method.
type UpdateGroupRequest struct {
	GroupID  string      `json:"group_id"`
	NewGroup interface{} `json:"new_group"`
}

// UpdateGroupResponse collects the response parameters for the UpdateGroup method.
type UpdateGroupResponse struct {
	Group interface{} `json:"group"`
	Err   error       `json:"err"`
}

// MakeUpdateGroupEndpoint returns an endpoint that invokes UpdateGroup on the service.
func MakeUpdateGroupEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateGroupRequest)
		group, err := s.UpdateGroup(ctx, req.GroupID, req.NewGroup)
		return UpdateGroupResponse{
			Err:   err,
			Group: group,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateGroupResponse) Failed() error {
	return r.Err
}

// GetGroupByIDRequest collects the request parameters for the GetGroupByID method.
type GetGroupByIDRequest struct {
	GroupID string `json:"group_id"`
}

// GetGroupByIDResponse collects the response parameters for the GetGroupByID method.
type GetGroupByIDResponse struct {
	Group interface{} `json:"group"`
	Err   error       `json:"err"`
}

// MakeGetGroupByIDEndpoint returns an endpoint that invokes GetGroupByID on the service.
func MakeGetGroupByIDEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGroupByIDRequest)
		group, err := s.GetGroupByID(ctx, req.GroupID)
		return GetGroupByIDResponse{
			Err:   err,
			Group: group,
		}, nil
	}
}

// Failed implements Failer.
func (r GetGroupByIDResponse) Failed() error {
	return r.Err
}

// GetGroupByNameRequest collects the request parameters for the GetGroupByName method.
type GetGroupByNameRequest struct {
	GroupName string `json:"group_name"`
}

// GetGroupByNameResponse collects the response parameters for the GetGroupByName method.
type GetGroupByNameResponse struct {
	Group interface{} `json:"group"`
	Err   error       `json:"err"`
}

// MakeGetGroupByNameEndpoint returns an endpoint that invokes GetGroupByName on the service.
func MakeGetGroupByNameEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGroupByNameRequest)
		group, err := s.GetGroupByName(ctx, req.GroupName)
		return GetGroupByNameResponse{
			Err:   err,
			Group: group,
		}, nil
	}
}

// Failed implements Failer.
func (r GetGroupByNameResponse) Failed() error {
	return r.Err
}

// IsGroupPrivateRequest collects the request parameters for the IsGroupPrivate method.
type IsGroupPrivateRequest struct {
	GroupID string `json:"group_id"`
}

// IsGroupPrivateResponse collects the response parameters for the IsGroupPrivate method.
type IsGroupPrivateResponse struct {
	IsPrivate bool  `json:"is_private"`
	Err       error `json:"err"`
}

// MakeIsGroupPrivateEndpoint returns an endpoint that invokes IsGroupPrivate on the service.
func MakeIsGroupPrivateEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(IsGroupPrivateRequest)
		isPrivate, err := s.IsGroupPrivate(ctx, req.GroupID)
		return IsGroupPrivateResponse{
			Err:       err,
			IsPrivate: isPrivate,
		}, nil
	}
}

// Failed implements Failer.
func (r IsGroupPrivateResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateUserAccount implements Service. Primarily useful in a client.
func (e Endpoints) CreateUserAccount(ctx context.Context, user interface{}, userType string) (e0 error) {
	request := CreateUserAccountRequest{
		User:     user,
		UserType: userType,
	}
	response, err := e.CreateUserAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateUserAccountResponse).E0
}

// UpdateUserAccount implements Service. Primarily useful in a client.
func (e Endpoints) UpdateUserAccount(ctx context.Context, user interface{}, id string) (e0 error) {
	request := UpdateUserAccountRequest{
		Id:   id,
		User: user,
	}
	response, err := e.UpdateUserAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateUserAccountResponse).E0
}

// DeleteUserAccount implements Service. Primarily useful in a client.
func (e Endpoints) DeleteUserAccount(ctx context.Context, id string) (e0 error) {
	request := DeleteUserAccountRequest{Id: id}
	response, err := e.DeleteUserAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteUserAccountResponse).E0
}

// GetUserAccount implements Service. Primarily useful in a client.
func (e Endpoints) GetUserAccount(ctx context.Context, id string) (user interface{}, err error) {
	request := GetUserAccountRequest{Id: id}
	response, err := e.GetUserAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserAccountResponse).User, response.(GetUserAccountResponse).Err
}

// CreateTeamsAccount implements Service. Primarily useful in a client.
func (e Endpoints) CreateTeamsAccount(ctx context.Context, team interface{}, teamType string) (e0 error) {
	request := CreateTeamsAccountRequest{
		Team:     team,
		TeamType: teamType,
	}
	response, err := e.CreateTeamsAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateTeamsAccountResponse).E0
}

// UpdateTeamsAccount implements Service. Primarily useful in a client.
func (e Endpoints) UpdateTeamsAccount(ctx context.Context, team interface{}, teamID string) (e0 error) {
	request := UpdateTeamsAccountRequest{
		Team:   team,
		TeamID: teamID,
	}
	response, err := e.UpdateTeamsAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateTeamsAccountResponse).E0
}

// DeleteTeamsAccount implements Service. Primarily useful in a client.
func (e Endpoints) DeleteTeamsAccount(ctx context.Context, teamID string) (e0 error) {
	request := DeleteTeamsAccountRequest{TeamID: teamID}
	response, err := e.DeleteTeamsAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteTeamsAccountResponse).E0
}

// GetTeamsAccount implements Service. Primarily useful in a client.
func (e Endpoints) GetTeamsAccount(ctx context.Context, teamID string) (team interface{}, err error) {
	request := GetTeamsAccountRequest{TeamID: teamID}
	response, err := e.GetTeamsAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTeamsAccountResponse).Team, response.(GetTeamsAccountResponse).Err
}

// AddUserToTeamsAccount implements Service. Primarily useful in a client.
func (e Endpoints) AddUserToTeamsAccount(ctx context.Context, teamID string, userID string) (e0 error) {
	request := AddUserToTeamsAccountRequest{
		TeamID: teamID,
		UserID: userID,
	}
	response, err := e.AddUserToTeamsAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddUserToTeamsAccountResponse).E0
}

// DeleteUserFromTeamsAccount implements Service. Primarily useful in a client.
func (e Endpoints) DeleteUserFromTeamsAccount(ctx context.Context, teamID string, userID string) (e0 error) {
	request := DeleteUserFromTeamsAccountRequest{
		TeamID: teamID,
		UserID: userID,
	}
	response, err := e.DeleteUserFromTeamsAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteUserFromTeamsAccountResponse).E0
}

// CreateGroup implements Service. Primarily useful in a client.
func (e Endpoints) CreateGroup(ctx context.Context, groupName string, groupType string, isPrivate bool) (e0 error) {
	request := CreateGroupRequest{
		GroupName: groupName,
		GroupType: groupType,
		IsPrivate: isPrivate,
	}
	response, err := e.CreateGroupEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateGroupResponse).E0
}

// SubscribeToGroup implements Service. Primarily useful in a client.
func (e Endpoints) SubscribeToGroup(ctx context.Context, groupID string, ID string) (e0 error) {
	request := SubscribeToGroupRequest{
		GroupID: groupID,
		ID:      ID,
	}
	response, err := e.SubscribeToGroupEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SubscribeToGroupResponse).E0
}

// UnsubscribeFromGroup implements Service. Primarily useful in a client.
func (e Endpoints) UnsubscribeFromGroup(ctx context.Context, groupID string, ID string) (e0 error) {
	request := UnsubscribeFromGroupRequest{
		GroupID: groupID,
		ID:      ID,
	}
	response, err := e.UnsubscribeFromGroupEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UnsubscribeFromGroupResponse).E0
}

// DeleteGroup implements Service. Primarily useful in a client.
func (e Endpoints) DeleteGroup(ctx context.Context, groupID string, groupAdminID string) (e0 error) {
	request := DeleteGroupRequest{
		GroupAdminID: groupAdminID,
		GroupID:      groupID,
	}
	response, err := e.DeleteGroupEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteGroupResponse).E0
}

// UpdateGroup implements Service. Primarily useful in a client.
func (e Endpoints) UpdateGroup(ctx context.Context, groupID string, newGroup interface{}) (group interface{}, err error) {
	request := UpdateGroupRequest{
		GroupID:  groupID,
		NewGroup: newGroup,
	}
	response, err := e.UpdateGroupEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateGroupResponse).Group, response.(UpdateGroupResponse).Err
}

// GetGroupByID implements Service. Primarily useful in a client.
func (e Endpoints) GetGroupByID(ctx context.Context, groupID string) (group interface{}, err error) {
	request := GetGroupByIDRequest{GroupID: groupID}
	response, err := e.GetGroupByIDEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetGroupByIDResponse).Group, response.(GetGroupByIDResponse).Err
}

// GetGroupByName implements Service. Primarily useful in a client.
func (e Endpoints) GetGroupByName(ctx context.Context, groupName string) (group interface{}, err error) {
	request := GetGroupByNameRequest{GroupName: groupName}
	response, err := e.GetGroupByNameEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetGroupByNameResponse).Group, response.(GetGroupByNameResponse).Err
}

// IsGroupPrivate implements Service. Primarily useful in a client.
func (e Endpoints) IsGroupPrivate(ctx context.Context, groupID string) (isPrivate bool, err error) {
	request := IsGroupPrivateRequest{GroupID: groupID}
	response, err := e.IsGroupPrivateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(IsGroupPrivateResponse).IsPrivate, response.(IsGroupPrivateResponse).Err
}
