package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(SocialService) SocialService

type loggingMiddleware struct {
	logger log.Logger
	next   SocialService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a SocialService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next SocialService) SocialService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) FollowUser(ctx context.Context, followerID string, followedID string) (e0 error) {
	defer func() {
		l.logger.Log("method", "FollowUser", "followerID", followerID, "followedID", followedID, "e0", e0)
	}()
	return l.next.FollowUser(ctx, followerID, followedID)
}
func (l loggingMiddleware) UnFollowerUser(ctx context.Context, followerID string, followedID string) (e0 error) {
	defer func() {
		l.logger.Log("method", "UnFollowerUser", "followerID", followerID, "followedID", followedID, "e0", e0)
	}()
	return l.next.UnFollowerUser(ctx, followerID, followedID)
}
func (l loggingMiddleware) FollowGroup(ctx context.Context, followerID string, groupID string) (e0 error) {
	defer func() {
		l.logger.Log("method", "FollowGroup", "followerID", followerID, "groupID", groupID, "e0", e0)
	}()
	return l.next.FollowGroup(ctx, followerID, groupID)
}
func (l loggingMiddleware) UnFollowGroup(ctx context.Context, followerID string, groupID string) (e0 error) {
	defer func() {
		l.logger.Log("method", "UnFollowGroup", "followerID", followerID, "groupID", groupID, "e0", e0)
	}()
	return l.next.UnFollowGroup(ctx, followerID, groupID)
}
func (l loggingMiddleware) FollowTeam(ctx context.Context, followerID string, teamID string) (e0 error) {
	defer func() {
		l.logger.Log("method", "FollowTeam", "followerID", followerID, "teamID", teamID, "e0", e0)
	}()
	return l.next.FollowTeam(ctx, followerID, teamID)
}
func (l loggingMiddleware) UnFollowTeam(ctx context.Context, followerID string, teamID string) (e0 error) {
	defer func() {
		l.logger.Log("method", "UnFollowTeam", "followerID", followerID, "teamID", teamID, "e0", e0)
	}()
	return l.next.UnFollowTeam(ctx, followerID, teamID)
}
