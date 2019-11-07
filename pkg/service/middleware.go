package service

import (
	"context"

	"github.com/go-kit/kit/metrics"
	"go.uber.org/zap"
)

// Middleware describes a service specific middleware
type Middleware func(Service) Service

// LoggingMiddleware returns a logging middleware that logs
// request specific information. From input arguments, to errors
func LoggingMiddleware(logger *zap.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{logger, next}
	}
}

// Logging Struct that implements the Service interface
type loggingMiddleware struct {
	logger *zap.Logger
	next Service
}

func (mw loggingMiddleware) GetUserById(ctx context.Context, id string) (user User, err error) {
	defer func(){
		if err != nil {
			mw.logger.Info("Request Completed",
				zap.String("method", "GetUserById"),
				zap.Any("user", user), zap.Any("error", err))
		}
	}()

	user, err = mw.next.GetUserById(ctx, id)

	if err != nil {
		return User{},err
	}
	return user,nil
}

func (mw loggingMiddleware) GetUserByEmail(ctx context.Context, email string) (user User, err error) {
	defer func(){
		if err != nil {
			mw.logger.Info("Request Completed",
				zap.String("method", "GetUserByEmail"),
				zap.Any("user", user), zap.Any("error", err))
		}
	}()

	user, err = mw.next.GetUserByEmail(ctx, email)

	if err != nil {
		return User{},err
	}
	return user,nil
}

func (mw loggingMiddleware) GetUserByUsername(ctx context.Context, username string) (user User, err error) {
	defer func(){
		if err != nil {
			mw.logger.Info("Request Completed",
				zap.String("method", "GetUserByUsername"),
				zap.Any("user", user), zap.Any("error", err))
		}
	}()

	user, err = mw.next.GetUserByUsername(ctx, username)

	if err != nil {
		return User{},err
	}
	return user,nil
}

// A logging wrapper around the create user service implementation
func (mw loggingMiddleware) CreateUser(ctx context.Context, user User) (err error) {
	defer func(){
		if err != nil {
			mw.logger.Info("Request Completed",
				zap.String("method", "CreateUser"),
				zap.Any("user", user), zap.Any("error", err))
		}
	}()

	err = mw.next.CreateUser(ctx, user)

	if err != nil {
		return err
	}
	return nil
}


// InstrumentingMiddleware returns a service middleware that instruments
// the number of users created over the lifetime of
// the service.
func InstrumentingMiddleware(CreateUserRequest, successfulCreateUserReq,
	failedCreateUserReq, getUserRequests, successfulGetUserReq, failedGetUserReq  metrics.Counter) Middleware {
	return func(next Service) Service {
		return instrumentingMiddleware{
			UsersCreateRequests:  CreateUserRequest,
			FailedUserCreateRequests: failedCreateUserReq,
			SuccessfulUserCreateRequests: successfulCreateUserReq,
			SuccessfulGetUserRequests:successfulGetUserReq,
			FailedGetUserRequests:failedGetUserReq,
			GetUserRequests: getUserRequests,
			next:  next,
		}
	}
}

// Instrumentation struct that implements the Service interface
type instrumentingMiddleware struct {
	UsersCreateRequests  metrics.Counter
	FailedUserCreateRequests metrics.Counter
	SuccessfulUserCreateRequests metrics.Counter
	SuccessfulGetUserRequests metrics.Counter
	FailedGetUserRequests metrics.Counter
	GetUserRequests metrics.Counter
	next  Service
}

func (mw instrumentingMiddleware) GetUserById(ctx context.Context, id string) (user User, err error) {
	mw.GetUserRequests.Add(1)
	user, err = mw.next.GetUserById(ctx, id)

	if err != nil {
		mw.FailedGetUserRequests.Add(1)
		return User{},err
	}

	mw.SuccessfulGetUserRequests.Add(1)
	return user,nil
}

func (mw instrumentingMiddleware) GetUserByEmail(ctx context.Context, email string) (user User, err error) {
	mw.GetUserRequests.Add(1)
	user, err = mw.next.GetUserByEmail(ctx, email)

	if err != nil {
		mw.FailedGetUserRequests.Add(1)
		return User{},err
	}

	mw.SuccessfulGetUserRequests.Add(1)
	return user,nil
}

func (mw instrumentingMiddleware) GetUserByUsername(ctx context.Context, username string) (user User, err error) {
	mw.GetUserRequests.Add(1)
	user, err = mw.next.GetUserByUsername(ctx, username)

	if err != nil {
		mw.FailedGetUserRequests.Add(1)
		return User{},err
	}

	mw.SuccessfulGetUserRequests.Add(1)
	return user,nil
}

// An instrumenting wrapper around the create user service implementation
func (mw instrumentingMiddleware) CreateUser(ctx context.Context, user User) (err error) {
	mw.UsersCreateRequests.Add(1)
	err = mw.next.CreateUser(ctx, user)

	if err != nil {
		mw.FailedUserCreateRequests.Add(1)
		return err
	}

	mw.SuccessfulUserCreateRequests.Add(1)
	return nil
}
