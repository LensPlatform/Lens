package service

import "context"

// SocialService describes the service.
type SocialService interface {
	// These interface definitions are particular to the social service
	FollowUser(ctx context.Context, followerID string, followedID string) error
	UnFollowerUser(ctx context.Context, followerID string, followedID string) error
	FollowGroup(ctx context.Context, followerID string, groupID string) error
	UnFollowGroup(ctx context.Context, followerID string, groupID string) error
	FollowTeam(ctx context.Context, followerID string, teamID string) error
	UnFollowTeam(ctx context.Context, followerID string, teamID string) error
}

type basicSocialService struct{}

func (b *basicSocialService) FollowUser(ctx context.Context, followerID string, followedID string) (e0 error) {
	// TODO implement the business logic of FollowUser
	return e0
}
func (b *basicSocialService) UnFollowerUser(ctx context.Context, followerID string, followedID string) (e0 error) {
	// TODO implement the business logic of UnFollowerUser
	return e0
}
func (b *basicSocialService) FollowGroup(ctx context.Context, followerID string, groupID string) (e0 error) {
	// TODO implement the business logic of FollowGroup
	return e0
}
func (b *basicSocialService) UnFollowGroup(ctx context.Context, followerID string, groupID string) (e0 error) {
	// TODO implement the business logic of UnFollowGroup
	return e0
}
func (b *basicSocialService) FollowTeam(ctx context.Context, followerID string, teamID string) (e0 error) {
	// TODO implement the business logic of FollowTeam
	return e0
}
func (b *basicSocialService) UnFollowTeam(ctx context.Context, followerID string, teamID string) (e0 error) {
	// TODO implement the business logic of UnFollowTeam
	return e0
}

// NewBasicSocialService returns a naive, stateless implementation of SocialService.
func NewBasicSocialService() SocialService {
	return &basicSocialService{}
}

// New returns a SocialService with all of the expected middleware wired in.
func New(middleware []Middleware) SocialService {
	var svc SocialService = NewBasicSocialService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
