package service

import (
	"context"
	"errors"

	sdetcd "github.com/go-kit/kit/sd/etcd"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type basicUsersService struct{}

type ConnectionProperties struct {
	maxRetryCount int
}

type UsersService interface {
	// User accounts specific interfaces. It is important to note that
	// users can be of three types: investor, startup member, and regular user.
	// Startup members are automatically associated with a team/startup. Investor
	// type users can be standalone or associated with an investor team account.
	// Regular users are not tied to any team and are inherently standalone.
	CreateUserAccount(ctx context.Context, user interface{}, userType string) (string, error)
	UpdateUserAccount(ctx context.Context, user interface{}, id string) (string, error)
	DeleteUserAccount(ctx context.Context, id string) (string, error)
	GetUserAccount(ctx context.Context, id string) (user interface{}, err error)

	// This presents interfaces tied to team accounts. Team accounts can be of two types:
	// startup teams or investor teams. Such teams are comprised of different user types and
	// fields detailing data specific to the team function.
	CreateTeamsAccount(ctx context.Context, team interface{}, teamType string) (string, error)
	UpdateTeamsAccount(ctx context.Context, team interface{}, teamID string) (string, error)
	DeleteTeamsAccount(ctx context.Context, teamID string) (string, error)
	GetTeamsAccount(ctx context.Context, teamID string) (team interface{}, err error)
	AddUserToTeamsAccount(ctx context.Context, teamID string, userID string) (string, error)
	DeleteUserFromTeamsAccount(ctx context.Context, teamID string, userID string) (string, error)

	// This presents interfaces tied to the group mechanism. Users can create user groups/communities
	// and post content within such communities. Additionally, groups can be public or private
	// hence, user can subscribe and unsubscribe to content posted in such a group.
	CreateGroup(ctx context.Context, groupName string, groupType string, isPrivate bool) (string, error)
	SubscribeToGroup(ctx context.Context, groupID string, ID string /*teams or user*/) (string, error)
	UnsubscribeFromGroup(ctx context.Context, groupID string, ID string /*teams or user*/) (string, error)
	DeleteGroup(ctx context.Context, groupID string, groupAdminID string) (string, error)
	UpdateGroup(ctx context.Context, groupID string, newGroup interface{}) (group interface{}, err error)
	GetGroupByID(ctx context.Context, groupID string) (group interface{}, err error)
	GetGroupByName(ctx context.Context, groupName string) (group interface{}, err error)
	IsGroupPrivate(ctx context.Context, groupID string) (isPrivate bool, err error)
}

func (b *basicUsersService) CreateUserAccount(ctx context.Context, user interface{}, userType string) (id string, e0 error) {
	// TODO implement the business logic of UpdateUserAccount
	return "", e0
}

func (b *basicUsersService) UpdateUserAccount(ctx context.Context, user interface{}, id string) (userID string, e0 error) {
	// TODO implement the business logic of UpdateUserAccount
	return "", e0
}
func (b *basicUsersService) DeleteUserAccount(ctx context.Context, id string) (userID string, e0 error) {
	// TODO implement the business logic of DeleteUserAccount
	return "", e0
}
func (b *basicUsersService) GetUserAccount(ctx context.Context, id string) (user interface{}, err error) {
	// TODO implement the business logic of GetUserAccount
	return user, err
}
func (b *basicUsersService) CreateTeamsAccount(ctx context.Context, team interface{}, teamType string) (id string, e0 error) {
	// TODO implement the business logic of CreateTeamsAccount
	return "", e0
}
func (b *basicUsersService) UpdateTeamsAccount(ctx context.Context, team interface{}, teamID string) (id string, e0 error) {
	// TODO implement the business logic of UpdateTeamsAccount
	return "", e0
}
func (b *basicUsersService) DeleteTeamsAccount(ctx context.Context, teamID string) (id string, e0 error) {
	// TODO implement the business logic of DeleteTeamsAccount
	return "", e0
}
func (b *basicUsersService) GetTeamsAccount(ctx context.Context, teamID string) (team interface{}, err error) {
	// TODO implement the business logic of GetTeamsAccount
	return team, err
}
func (b *basicUsersService) AddUserToTeamsAccount(ctx context.Context, teamID string, userID string) (id string, e0 error) {
	// TODO implement the business logic of AddUserToTeamsAccount
	return "", e0
}
func (b *basicUsersService) DeleteUserFromTeamsAccount(ctx context.Context, teamID string, userID string) (id string, e0 error) {
	// TODO implement the business logic of DeleteUserFromTeamsAccount
	return "", e0
}
func (b *basicUsersService) CreateGroup(ctx context.Context, groupName string, groupType string, isPrivate bool) (id string, e0 error) {
	// TODO implement the business logic of CreateGroup
	return "", e0
}
func (b *basicUsersService) SubscribeToGroup(ctx context.Context, groupID string, ID string) (id string, e0 error) {
	// TODO implement the business logic of SubscribeToGroup
	return "", e0
}
func (b *basicUsersService) UnsubscribeFromGroup(ctx context.Context, groupID string, ID string) (id string, e0 error) {
	// TODO implement the business logic of UnsubscribeFromGroup
	return "", e0
}
func (b *basicUsersService) DeleteGroup(ctx context.Context, groupID string, groupAdminID string) (id string, e0 error) {
	// TODO implement the business logic of DeleteGroup
	return "", e0
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

func connectToServices(Logger *zap.Logger, connectionProperties ConnectionProperties, servicePrefixes ...string /* ex. "/services/notificator/"*/) error {
	var etcdServer = "http://etcd:2379"
	client, err := sdetcd.NewClient(context.Background(), []string{etcdServer}, sdetcd.ClientOptions{})
	if err != nil {
		Logger.Error("unable to connect to etcd: %s", zap.Error(err))
		// TODO: Return option must be present
		return err
	}

	// Make sure service Prefix slice is not empty and attempt connection to all entries
	// present in the slice
	for _, servicePrefix := range servicePrefixes {
		// Obtain number of service instances presently defined in the docker config.
		entries, err := client.GetEntries(servicePrefix)

		// perform error checking.
		if err != nil || len(entries) == 0 {
			Logger.Error("unable to get prefix entries: %s", zap.Error(err))
			return err
		}

		// connect to a service instance
		_, err = grpc.Dial(entries[0], grpc.WithInsecure())

		if err != nil {
			Logger.Error("unable to connect to notificator: %s. attempting retry", zap.Error(err))

			err = retryConnection(entries, 0, connectionProperties.maxRetryCount, Logger)

			if err != nil {
				Logger.Error("retry count exceeded. unable to connect to service", zap.Error(err))
				return err
			}
			return err
		}

		Logger.Info("successfully connected to service", zap.String("ServiceName", servicePrefix))
	}
	return nil
}

func retryConnection(serviceInstances []string, selectedInstanceIdx int, maxRetryCount int, Logger *zap.Logger) error {
	retryCount := 0
	lastInstanceIdx := len(serviceInstances) - 1
	for retryCount <= maxRetryCount {
		// attempt connection retry to an instances starting from the last instance
		_, err := grpc.Dial(serviceInstances[lastInstanceIdx], grpc.WithInsecure())

		lastInstanceIdx--

		if lastInstanceIdx != -1 {
			lastInstanceIdx = len(serviceInstances) - 1
		}

		// add error handling check.
		if err != nil {
			Logger.Error("unable to connect to notificator: %s", zap.Error(err))
			retryCount++
		} else {
			return nil
		}

	}

	return errors.New("unable to connect to service due to exceeded retry count")
}
