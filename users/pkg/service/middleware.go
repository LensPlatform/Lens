package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

type Middleware func(UsersService) UsersService

type loggingMiddleware struct {
	logger log.Logger
	next   UsersService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UsersService) UsersService {
		return &loggingMiddleware{logger, next}
	}
}

func (l loggingMiddleware) CreateUserAccount(ctx context.Context, user interface{}, userType string) (e0 error) {
	defer func() {
		l.logger.Log("method", "CreateUserAccount", "user", user, "userType", userType, "e0", e0)
	}()
	return l.next.CreateUserAccount(ctx, user, userType)
}
func (l loggingMiddleware) UpdateUserAccount(ctx context.Context, user interface{}, id string) (e0 error) {
	defer func() {
		l.logger.Log("method", "UpdateUserAccount", "user", user, "id", id, "e0", e0)
	}()
	return l.next.UpdateUserAccount(ctx, user, id)
}
func (l loggingMiddleware) DeleteUserAccount(ctx context.Context, id string) (e0 error) {
	defer func() {
		l.logger.Log("method", "DeleteUserAccount", "id", id, "e0", e0)
	}()
	return l.next.DeleteUserAccount(ctx, id)
}
func (l loggingMiddleware) GetUserAccount(ctx context.Context, id string) (user interface{}, err error) {
	defer func() {
		l.logger.Log("method", "GetUserAccount", "id", id, "user", user, "err", err)
	}()
	return l.next.GetUserAccount(ctx, id)
}
func (l loggingMiddleware) CreateTeamsAccount(ctx context.Context, team interface{}, teamType string) (e0 error) {
	defer func() {
		l.logger.Log("method", "CreateTeamsAccount", "team", team, "teamType", teamType, "e0", e0)
	}()
	return l.next.CreateTeamsAccount(ctx, team, teamType)
}
func (l loggingMiddleware) UpdateTeamsAccount(ctx context.Context, team interface{}, teamID string) (e0 error) {
	defer func() {
		l.logger.Log("method", "UpdateTeamsAccount", "team", team, "teamID", teamID, "e0", e0)
	}()
	return l.next.UpdateTeamsAccount(ctx, team, teamID)
}
func (l loggingMiddleware) DeleteTeamsAccount(ctx context.Context, teamID string) (e0 error) {
	defer func() {
		l.logger.Log("method", "DeleteTeamsAccount", "teamID", teamID, "e0", e0)
	}()
	return l.next.DeleteTeamsAccount(ctx, teamID)
}
func (l loggingMiddleware) GetTeamsAccount(ctx context.Context, teamID string) (team interface{}, err error) {
	defer func() {
		l.logger.Log("method", "GetTeamsAccount", "teamID", teamID, "team", team, "err", err)
	}()
	return l.next.GetTeamsAccount(ctx, teamID)
}
func (l loggingMiddleware) AddUserToTeamsAccount(ctx context.Context, teamID string, userID string) (e0 error) {
	defer func() {
		l.logger.Log("method", "AddUserToTeamsAccount", "teamID", teamID, "userID", userID, "e0", e0)
	}()
	return l.next.AddUserToTeamsAccount(ctx, teamID, userID)
}
func (l loggingMiddleware) DeleteUserFromTeamsAccount(ctx context.Context, teamID string, userID string) (e0 error) {
	defer func() {
		l.logger.Log("method", "DeleteUserFromTeamsAccount", "teamID", teamID, "userID", userID, "e0", e0)
	}()
	return l.next.DeleteUserFromTeamsAccount(ctx, teamID, userID)
}
func (l loggingMiddleware) CreateGroup(ctx context.Context, groupName string, groupType string, isPrivate bool) (e0 error) {
	defer func() {
		l.logger.Log("method", "CreateGroup", "groupName", groupName, "groupType", groupType, "isPrivate", isPrivate, "e0", e0)
	}()
	return l.next.CreateGroup(ctx, groupName, groupType, isPrivate)
}
func (l loggingMiddleware) SubscribeToGroup(ctx context.Context, groupID string, ID string) (e0 error) {
	defer func() {
		l.logger.Log("method", "SubscribeToGroup", "groupID", groupID, "ID", ID, "e0", e0)
	}()
	return l.next.SubscribeToGroup(ctx, groupID, ID)
}
func (l loggingMiddleware) UnsubscribeFromGroup(ctx context.Context, groupID string, ID string) (e0 error) {
	defer func() {
		l.logger.Log("method", "UnsubscribeFromGroup", "groupID", groupID, "ID", ID, "e0", e0)
	}()
	return l.next.UnsubscribeFromGroup(ctx, groupID, ID)
}
func (l loggingMiddleware) DeleteGroup(ctx context.Context, groupID string, groupAdminID string) (e0 error) {
	defer func() {
		l.logger.Log("method", "DeleteGroup", "groupID", groupID, "groupAdminID", groupAdminID, "e0", e0)
	}()
	return l.next.DeleteGroup(ctx, groupID, groupAdminID)
}
func (l loggingMiddleware) UpdateGroup(ctx context.Context, groupID string, newGroup interface{}) (group interface{}, err error) {
	defer func() {
		l.logger.Log("method", "UpdateGroup", "groupID", groupID, "newGroup", newGroup, "group", group, "err", err)
	}()
	return l.next.UpdateGroup(ctx, groupID, newGroup)
}
func (l loggingMiddleware) GetGroupByID(ctx context.Context, groupID string) (group interface{}, err error) {
	defer func() {
		l.logger.Log("method", "GetGroupByID", "groupID", groupID, "group", group, "err", err)
	}()
	return l.next.GetGroupByID(ctx, groupID)
}
func (l loggingMiddleware) GetGroupByName(ctx context.Context, groupName string) (group interface{}, err error) {
	defer func() {
		l.logger.Log("method", "GetGroupByName", "groupName", groupName, "group", group, "err", err)
	}()
	return l.next.GetGroupByName(ctx, groupName)
}
func (l loggingMiddleware) IsGroupPrivate(ctx context.Context, groupID string) (isPrivate bool, err error) {
	defer func() {
		l.logger.Log("method", "IsGroupPrivate", "groupID", groupID, "isPrivate", isPrivate, "err", err)
	}()
	return l.next.IsGroupPrivate(ctx, groupID)
}
