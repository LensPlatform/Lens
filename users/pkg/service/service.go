package service

import (
	"context"
)

// UsersService describes the service.
type UsersService interface {
	// User accounts specific interfaces. It is important to note that
	// users can be of three types: investor, startup member, and regular user.
	// Startup members are automatically associated with a team/startup. Investor
	// type users can be standalone or associated with an investor team account.
	// Regular users are not tied to any team and are inherently standalone.
	CreateUserAccount(ctx context.Context, user interface{}, userType string) error
	UpdateUserAccount(ctx context.Context, user interface{}, id string) error
	DeleteUserAccount(ctx context.Context, id string) error
	GetUserAccount(ctx context.Context, id string) (user interface{}, err error)

	// This presents interfaces tied to team accounts. Team accounts can be of two types:
	// startup teams or investor teams. Such teams are comprised of different user types and
	// fields detailing data specific to the team function.
	CreateTeamsAccount(ctx context.Context, team interface{}, teamType string) error
	UpdateTeamsAccount(ctx context.Context, team interface{}, teamID string) error
	DeleteTeamsAccount(ctx context.Context, teamID string) error
	GetTeamsAccount(ctx context.Context, teamID string) (team interface{}, err error)
	AddUserToTeamsAccount(ctx context.Context, teamID string, userID string) error
	DeleteUserFromTeamsAccount(ctx context.Context, teamID string, userID string) error

	// This presents interfaces tied to the group mechanism. Users can create user groups/communities
	// and post content within such communities. Additionally, groups can be public or private
	// hence, user can subscribe and unsubscribe to content posted in such a group.
	CreateGroup(ctx context.Context, groupName string, groupType string, isPrivate bool) error
	SubscribeToGroup(ctx context.Context, groupID string, ID string /*teams or user*/) error
	UnsubscribeFromGroup(ctx context.Context, groupID string, ID string /*teams or user*/) error
	DeleteGroup(ctx context.Context, groupID string, groupAdminID string) error
	UpdateGroup(ctx context.Context, groupID string, newGroup interface{}) (group interface{}, err error)
	GetGroupByID(ctx context.Context, groupID string) (group interface{}, err error)
	GetGroupByName(ctx context.Context, groupName string) (group interface{}, err error)
	IsGroupPrivate(ctx context.Context, groupID string) (isPrivate bool, err error)
}

type basicUsersService struct{}

func (b *basicUsersService) CreateUserAccount(ctx context.Context, user interface{}, userType string) (e0 error) {
	// TODO implement the business logic of CreateUserAccount
	return e0
}
func (b *basicUsersService) UpdateUserAccount(ctx context.Context, user interface{}, id string) (e0 error) {
	// TODO implement the business logic of UpdateUserAccount
	return e0
}
func (b *basicUsersService) DeleteUserAccount(ctx context.Context, id string) (e0 error) {
	// TODO implement the business logic of DeleteUserAccount
	return e0
}
func (b *basicUsersService) GetUserAccount(ctx context.Context, id string) (user interface{}, err error) {
	// TODO implement the business logic of GetUserAccount
	return user, err
}
func (b *basicUsersService) CreateTeamsAccount(ctx context.Context, team interface{}, teamType string) (e0 error) {
	// TODO implement the business logic of CreateTeamsAccount
	return e0
}
func (b *basicUsersService) UpdateTeamsAccount(ctx context.Context, team interface{}, teamID string) (e0 error) {
	// TODO implement the business logic of UpdateTeamsAccount
	return e0
}
func (b *basicUsersService) DeleteTeamsAccount(ctx context.Context, teamID string) (e0 error) {
	// TODO implement the business logic of DeleteTeamsAccount
	return e0
}
func (b *basicUsersService) GetTeamsAccount(ctx context.Context, teamID string) (team interface{}, err error) {
	// TODO implement the business logic of GetTeamsAccount
	return team, err
}
func (b *basicUsersService) AddUserToTeamsAccount(ctx context.Context, teamID string, userID string) (e0 error) {
	// TODO implement the business logic of AddUserToTeamsAccount
	return e0
}
func (b *basicUsersService) DeleteUserFromTeamsAccount(ctx context.Context, teamID string, userID string) (e0 error) {
	// TODO implement the business logic of DeleteUserFromTeamsAccount
	return e0
}
func (b *basicUsersService) CreateGroup(ctx context.Context, groupName string, groupType string, isPrivate bool) (e0 error) {
	// TODO implement the business logic of CreateGroup
	return e0
}
func (b *basicUsersService) SubscribeToGroup(ctx context.Context, groupID string, ID string) (e0 error) {
	// TODO implement the business logic of SubscribeToGroup
	return e0
}
func (b *basicUsersService) UnsubscribeFromGroup(ctx context.Context, groupID string, ID string) (e0 error) {
	// TODO implement the business logic of UnsubscribeFromGroup
	return e0
}
func (b *basicUsersService) DeleteGroup(ctx context.Context, groupID string, groupAdminID string) (e0 error) {
	// TODO implement the business logic of DeleteGroup
	return e0
}
func (b *basicUsersService) UpdateGroup(ctx context.Context, groupID string, newGroup interface{}) (group interface{}, err error) {
	// TODO implement the business logic of UpdateGroup
	return group, err
}
func (b *basicUsersService) GetGroupByID(ctx context.Context, groupID string) (group interface{}, err error) {
	// TODO implement the business logic of GetGroupByID
	return group, err
}
func (b *basicUsersService) GetGroupByName(ctx context.Context, groupName string) (group interface{}, err error) {
	// TODO implement the business logic of GetGroupByName
	return group, err
}
func (b *basicUsersService) IsGroupPrivate(ctx context.Context, groupID string) (isPrivate bool, err error) {
	// TODO implement the business logic of IsGroupPrivate
	return isPrivate, err
}

// NewBasicUsersService returns a naive, stateless implementation of UsersService.
func NewBasicUsersService() UsersService {
	return &basicUsersService{}
}

// New returns a UsersService with all of the expected middleware wired in.
func New(middleware []Middleware) UsersService {
	var svc UsersService = NewBasicUsersService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
