package endpoint

import (
	"context"

	service "github.com/LensPlatform/Lens/social/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// FollowUserRequest collects the request parameters for the FollowUser method.
type FollowUserRequest struct {
	FollowerID string `json:"follower_id"`
	FollowedID string `json:"followed_id"`
}

// FollowUserResponse collects the response parameters for the FollowUser method.
type FollowUserResponse struct {
	E0 error `json:"e0"`
}

// MakeFollowUserEndpoint returns an endpoint that invokes FollowUser on the service.
func MakeFollowUserEndpoint(s service.SocialService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FollowUserRequest)
		e0 := s.FollowUser(ctx, req.FollowerID, req.FollowedID)
		return FollowUserResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r FollowUserResponse) Failed() error {
	return r.E0
}

// UnFollowerUserRequest collects the request parameters for the UnFollowerUser method.
type UnFollowerUserRequest struct {
	FollowerID string `json:"follower_id"`
	FollowedID string `json:"followed_id"`
}

// UnFollowerUserResponse collects the response parameters for the UnFollowerUser method.
type UnFollowerUserResponse struct {
	E0 error `json:"e0"`
}

// MakeUnFollowerUserEndpoint returns an endpoint that invokes UnFollowerUser on the service.
func MakeUnFollowerUserEndpoint(s service.SocialService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UnFollowerUserRequest)
		e0 := s.UnFollowerUser(ctx, req.FollowerID, req.FollowedID)
		return UnFollowerUserResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r UnFollowerUserResponse) Failed() error {
	return r.E0
}

// FollowGroupRequest collects the request parameters for the FollowGroup method.
type FollowGroupRequest struct {
	FollowerID string `json:"follower_id"`
	GroupID    string `json:"group_id"`
}

// FollowGroupResponse collects the response parameters for the FollowGroup method.
type FollowGroupResponse struct {
	E0 error `json:"e0"`
}

// MakeFollowGroupEndpoint returns an endpoint that invokes FollowGroup on the service.
func MakeFollowGroupEndpoint(s service.SocialService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FollowGroupRequest)
		e0 := s.FollowGroup(ctx, req.FollowerID, req.GroupID)
		return FollowGroupResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r FollowGroupResponse) Failed() error {
	return r.E0
}

// UnFollowGroupRequest collects the request parameters for the UnFollowGroup method.
type UnFollowGroupRequest struct {
	FollowerID string `json:"follower_id"`
	GroupID    string `json:"group_id"`
}

// UnFollowGroupResponse collects the response parameters for the UnFollowGroup method.
type UnFollowGroupResponse struct {
	E0 error `json:"e0"`
}

// MakeUnFollowGroupEndpoint returns an endpoint that invokes UnFollowGroup on the service.
func MakeUnFollowGroupEndpoint(s service.SocialService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UnFollowGroupRequest)
		e0 := s.UnFollowGroup(ctx, req.FollowerID, req.GroupID)
		return UnFollowGroupResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r UnFollowGroupResponse) Failed() error {
	return r.E0
}

// FollowTeamRequest collects the request parameters for the FollowTeam method.
type FollowTeamRequest struct {
	FollowerID string `json:"follower_id"`
	TeamID     string `json:"team_id"`
}

// FollowTeamResponse collects the response parameters for the FollowTeam method.
type FollowTeamResponse struct {
	E0 error `json:"e0"`
}

// MakeFollowTeamEndpoint returns an endpoint that invokes FollowTeam on the service.
func MakeFollowTeamEndpoint(s service.SocialService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FollowTeamRequest)
		e0 := s.FollowTeam(ctx, req.FollowerID, req.TeamID)
		return FollowTeamResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r FollowTeamResponse) Failed() error {
	return r.E0
}

// UnFollowTeamRequest collects the request parameters for the UnFollowTeam method.
type UnFollowTeamRequest struct {
	FollowerID string `json:"follower_id"`
	TeamID     string `json:"team_id"`
}

// UnFollowTeamResponse collects the response parameters for the UnFollowTeam method.
type UnFollowTeamResponse struct {
	E0 error `json:"e0"`
}

// MakeUnFollowTeamEndpoint returns an endpoint that invokes UnFollowTeam on the service.
func MakeUnFollowTeamEndpoint(s service.SocialService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UnFollowTeamRequest)
		e0 := s.UnFollowTeam(ctx, req.FollowerID, req.TeamID)
		return UnFollowTeamResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r UnFollowTeamResponse) Failed() error {
	return r.E0
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// FollowUser implements Service. Primarily useful in a client.
func (e Endpoints) FollowUser(ctx context.Context, followerID string, followedID string) (e0 error) {
	request := FollowUserRequest{
		FollowedID: followedID,
		FollowerID: followerID,
	}
	response, err := e.FollowUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(FollowUserResponse).E0
}

// UnFollowerUser implements Service. Primarily useful in a client.
func (e Endpoints) UnFollowerUser(ctx context.Context, followerID string, followedID string) (e0 error) {
	request := UnFollowerUserRequest{
		FollowedID: followedID,
		FollowerID: followerID,
	}
	response, err := e.UnFollowerUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UnFollowerUserResponse).E0
}

// FollowGroup implements Service. Primarily useful in a client.
func (e Endpoints) FollowGroup(ctx context.Context, followerID string, groupID string) (e0 error) {
	request := FollowGroupRequest{
		FollowerID: followerID,
		GroupID:    groupID,
	}
	response, err := e.FollowGroupEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(FollowGroupResponse).E0
}

// UnFollowGroup implements Service. Primarily useful in a client.
func (e Endpoints) UnFollowGroup(ctx context.Context, followerID string, groupID string) (e0 error) {
	request := UnFollowGroupRequest{
		FollowerID: followerID,
		GroupID:    groupID,
	}
	response, err := e.UnFollowGroupEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UnFollowGroupResponse).E0
}

// FollowTeam implements Service. Primarily useful in a client.
func (e Endpoints) FollowTeam(ctx context.Context, followerID string, teamID string) (e0 error) {
	request := FollowTeamRequest{
		FollowerID: followerID,
		TeamID:     teamID,
	}
	response, err := e.FollowTeamEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(FollowTeamResponse).E0
}

// UnFollowTeam implements Service. Primarily useful in a client.
func (e Endpoints) UnFollowTeam(ctx context.Context, followerID string, teamID string) (e0 error) {
	request := UnFollowTeamRequest{
		FollowerID: followerID,
		TeamID:     teamID,
	}
	response, err := e.UnFollowTeamEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UnFollowTeamResponse).E0
}
