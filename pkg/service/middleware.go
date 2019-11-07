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
func InstrumentingMiddleware(request, success, failed metrics.Counter) Middleware {
	return func(next Service) Service {
		return instrumentingMiddleware{
			UsersCreateRequests:  request,
			FailedUserCreateRequests: failed,
			SuccessfulUserCreateRequests: success,
			next:  next,
		}
	}
}

// Instrumentation struct that implements the Service interface
type instrumentingMiddleware struct {
	UsersCreateRequests  metrics.Counter
	FailedUserCreateRequests metrics.Counter
	SuccessfulUserCreateRequests metrics.Counter
	next  Service
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
