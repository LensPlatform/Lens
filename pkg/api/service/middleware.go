package service

import (
	"context"
	"fmt"

	_ "github.com/go-kit/kit/log"
	"go.uber.org/zap"
)

type Middleware func(UsersService) UsersService

type loggingMiddleware struct {
	logger *zap.Logger
	next   UsersService
}

func LoggingMiddleware(logger *zap.Logger) Middleware {
	return func(next UsersService) UsersService {
		return &loggingMiddleware{logger, next}
	}
}

func (l loggingMiddleware) CreateUserAccount(ctx context.Context, user interface{}, userType string) (id string, e0 error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "CreateUserAccount"), zap.String("user", fmt.Sprintf("%v", user)),
			zap.String("userType", userType), zap.Any("err", e0))
	}()
	return l.next.CreateUserAccount(ctx, user, userType)
}
func (l loggingMiddleware) UpdateUserAccount(ctx context.Context, user interface{}, id string) (userID string, e0 error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "UpdateUserAccount"), zap.String("user", fmt.Sprintf("%v", user)),
			zap.String("id", id), zap.Any("err", e0))
	}()
	return l.next.UpdateUserAccount(ctx, user, id)
}
func (l loggingMiddleware) DeleteUserAccount(ctx context.Context, id string) (userId string, e0 error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "DeleteUserAccount"),
			zap.String("id", id), zap.Any("err", e0))
	}()
	return l.next.DeleteUserAccount(ctx, id)
}
func (l loggingMiddleware) GetUserAccount(ctx context.Context, id string) (user interface{}, err error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "GetUserAccount"), zap.String("user", fmt.Sprintf("%v", user)),
			zap.String("id", id), zap.Any("err", err))
	}()
	return l.next.GetUserAccount(ctx, id)
}
func (l loggingMiddleware) CreateTeamsAccount(ctx context.Context, team interface{}, teamType string) (id string, e0 error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "CreateTeamsAccount"), zap.String("team", fmt.Sprintf("%v", team)),
			zap.String("teamType", teamType), zap.Any("err", e0))
	}()
	return l.next.CreateTeamsAccount(ctx, team, teamType)
}
func (l loggingMiddleware) UpdateTeamsAccount(ctx context.Context, team interface{}, teamID string) (id string, e0 error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "UpdateTeamsAccount"), zap.String("team", fmt.Sprintf("%v", team)),
			zap.String("teamID", teamID), zap.Any("err", e0))
	}()
	return l.next.UpdateTeamsAccount(ctx, team, teamID)
}
func (l loggingMiddleware) DeleteTeamsAccount(ctx context.Context, teamID string) (id string, e0 error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "DeleteTeamsAccount"),
			zap.String("teamID", teamID), zap.Any("err", e0))
	}()
	return l.next.DeleteTeamsAccount(ctx, teamID)
}
func (l loggingMiddleware) GetTeamsAccount(ctx context.Context, teamID string) (team interface{}, err error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "GetTeamsAccount"), zap.String("team", fmt.Sprintf("%v", team)),
			zap.String("teamID", teamID), zap.Any("err", err))
	}()
	return l.next.GetTeamsAccount(ctx, teamID)
}
func (l loggingMiddleware) AddUserToTeamsAccount(ctx context.Context, teamID string, userID string) (id string, e0 error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "AddUserToTeamsAccount"), zap.String("userID", userID),
			zap.String("teamID", teamID), zap.Any("err", e0))
	}()
	return l.next.AddUserToTeamsAccount(ctx, teamID, userID)
}
func (l loggingMiddleware) DeleteUserFromTeamsAccount(ctx context.Context, teamID string, userID string) (id string, e0 error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "DeleteUserFromTeamsAccount"), zap.String("userID", userID),
			zap.String("teamID", teamID), zap.Any("err", e0))
	}()
	return l.next.DeleteUserFromTeamsAccount(ctx, teamID, userID)
}
func (l loggingMiddleware) CreateGroup(ctx context.Context, groupName string, groupType string, isPrivate bool) (id string, e0 error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "CreateGroup"), zap.String("groupName", groupName),
			zap.String("groupType", groupType), zap.Bool("isPrivate", isPrivate), zap.Any("err", e0))
	}()
	return l.next.CreateGroup(ctx, groupName, groupType, isPrivate)
}
func (l loggingMiddleware) SubscribeToGroup(ctx context.Context, groupID string, ID string) (id string, e0 error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "SubscribeToGroup"), zap.String("id", ID),
			zap.String("groupID", groupID), zap.Any("err", e0))
	}()
	return l.next.SubscribeToGroup(ctx, groupID, ID)
}
func (l loggingMiddleware) UnsubscribeFromGroup(ctx context.Context, groupID string, ID string) (id string, e0 error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "UnsubscribeFromGroup"), zap.String("id", ID),
			zap.String("groupID", groupID), zap.Any("err", e0))
	}()
	return l.next.UnsubscribeFromGroup(ctx, groupID, ID)
}
func (l loggingMiddleware) DeleteGroup(ctx context.Context, groupID string, groupAdminID string) (id string, e0 error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "DeleteGroup"), zap.String("groupAdminID", groupAdminID),
			zap.String("groupID", groupID), zap.Any("err", e0))
	}()
	return l.next.DeleteGroup(ctx, groupID, groupAdminID)
}
func (l loggingMiddleware) UpdateGroup(ctx context.Context, groupID string, newGroup interface{}) (group interface{}, err error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "UpdateGroup"),
			zap.String("groupID", groupID), zap.String("newGroup", fmt.Sprintf("%v", newGroup)),
			zap.String("Group", fmt.Sprintf("%v", group)),
			zap.Any("err", err))
	}()
	return l.next.UpdateGroup(ctx, groupID, newGroup)
}
func (l loggingMiddleware) GetGroupByID(ctx context.Context, groupID string) (group interface{}, err error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "GetGroupByID"),
			zap.String("group", fmt.Sprintf("%v", group)),
			zap.String("groupID", groupID), zap.Any("err", err))
	}()
	return l.next.GetGroupByID(ctx, groupID)
}
func (l loggingMiddleware) GetGroupByName(ctx context.Context, groupName string) (group interface{}, err error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "GetGroupByName"),
			zap.String("group", fmt.Sprintf("%v", group)),
			zap.String("groupName", groupName), zap.Any("err", err))
	}()
	return l.next.GetGroupByName(ctx, groupName)
}
func (l loggingMiddleware) IsGroupPrivate(ctx context.Context, groupID string) (isPrivate bool, err error) {
	defer func() {
		l.logger.Info("method", zap.String("function", "IsGroupPrivate"),
			zap.Bool("isPrivate", isPrivate),
			zap.String("groupID", groupID), zap.Any("err", err))
	}()
	return l.next.IsGroupPrivate(ctx, groupID)
}
