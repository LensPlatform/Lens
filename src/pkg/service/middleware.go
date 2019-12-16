package service

import (
	"context"

	"go.uber.org/zap"

	model "github.com/LensPlatform/Lens/src/pkg/models"
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
	next   Service
}

// A logging wrapper around the LogIn service implementation
func (mw loggingMiddleware) LogIn(ctx context.Context, username, password string) (user model.User, err error) {
	defer func() {
		if err != nil {
			mw.logger.Info("Request Completed",
				zap.String("method", "LogIn"),
				zap.Any("user", user), zap.Any("error", err))
		}
	}()

	user, err = mw.next.LogIn(ctx, username, password)

	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// A logging wrapper around the GetUserById service implementation
func (mw loggingMiddleware) GetUserById(ctx context.Context, id string) (user model.User, err error) {
	defer func() {
		if err != nil {
			mw.logger.Info("Request Completed",
				zap.String("method", "GetUserById"),
				zap.Any("user", user), zap.Any("error", err))
		}
	}()

	user, err = mw.next.GetUserById(ctx, id)

	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// A logging wrapper around the GetUserByEmail service implementation
func (mw loggingMiddleware) GetUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	defer func() {
		if err != nil {
			mw.logger.Info("Request Completed",
				zap.String("method", "GetUserByEmail"),
				zap.Any("user", user), zap.Any("error", err))
		}
	}()

	user, err = mw.next.GetUserByEmail(ctx, email)

	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// A logging wrapper around the GetUserByUsername service implementation
func (mw loggingMiddleware) GetUserByUsername(ctx context.Context, username string) (user model.User, err error) {
	defer func() {
		if err != nil {
			mw.logger.Info("Request Completed",
				zap.String("method", "GetUserByUsername"),
				zap.Any("user", user), zap.Any("error", err))
		}
	}()

	user, err = mw.next.GetUserByUsername(ctx, username)

	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// A logging wrapper around the create user service implementation
func (mw loggingMiddleware) CreateUser(ctx context.Context, user model.User) (err error) {
	defer func() {
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
func InstrumentingMiddleware(counters Counters) Middleware {
	return func(next Service) Service {
		var mw instrumentingMiddleware

		mw.CreateUserRequest = counters.CreateUserRequest
		mw.FailedCreateUserRequest = counters.FailedCreateUserRequest
		mw.SuccessfulCreateUserRequest = counters.SuccessfulCreateUserRequest
		mw.SuccessfulGetUserRequest = counters.SuccessfulGetUserRequest
		mw.FailedGetUserRequest = counters.FailedGetUserRequest
		mw.GetUserRequest = counters.GetUserRequest
		mw.SuccessfulLogInRequest = counters.SuccessfulLogInRequest
		mw.FailedLogInRequest = counters.FailedLogInRequest
		mw.next = next
		return mw
	}
}

// Instrumentation struct that implements the Service interface
type instrumentingMiddleware struct {
	Counters
	next Service
}

// An instrumenting wrapper around the LogIn service implementation
func (mw instrumentingMiddleware) LogIn(ctx context.Context, username, password string) (user model.User, err error) {
	user, err = mw.next.LogIn(ctx, username, password)

	if err != nil {
		mw.FailedLogInRequest.Add(1)
		return model.User{}, err
	}

	mw.SuccessfulLogInRequest.Add(1)
	return user, nil
}

// An instrumenting wrapper around the GetUserById service implementation
func (mw instrumentingMiddleware) GetUserById(ctx context.Context, id string) (user model.User, err error) {
	mw.GetUserRequest.Add(1)
	user, err = mw.next.GetUserById(ctx, id)

	if err != nil {
		mw.FailedGetUserRequest.Add(1)
		return model.User{}, err
	}

	mw.SuccessfulGetUserRequest.Add(1)
	return user, nil
}

// An instrumenting wrapper around the GetUserByEmail service implementation
func (mw instrumentingMiddleware) GetUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	mw.GetUserRequest.Add(1)
	user, err = mw.next.GetUserByEmail(ctx, email)

	if err != nil {
		mw.FailedGetUserRequest.Add(1)
		return model.User{}, err
	}

	mw.SuccessfulGetUserRequest.Add(1)
	return user, nil
}

// An instrumenting wrapper around the GetUserByUsername service implementation
func (mw instrumentingMiddleware) GetUserByUsername(ctx context.Context, username string) (user model.User, err error) {
	mw.GetUserRequest.Add(1)
	user, err = mw.next.GetUserByUsername(ctx, username)

	if err != nil {
		mw.FailedGetUserRequest.Add(1)
		return model.User{}, err
	}

	mw.SuccessfulGetUserRequest.Add(1)
	return user, nil
}

// An instrumenting wrapper around the create user service implementation
func (mw instrumentingMiddleware) CreateUser(ctx context.Context, user model.User) (err error) {
	mw.CreateUserRequest.Add(1)
	err = mw.next.CreateUser(ctx, user)

	if err != nil {
		mw.FailedCreateUserRequest.Add(1)
		return err
	}

	mw.SuccessfulCreateUserRequest.Add(1)
	return nil
}
