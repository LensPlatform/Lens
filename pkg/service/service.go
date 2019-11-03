package service

import (
	"context"

	"github.com/go-kit/kit/metrics"
	"go.uber.org/zap"
)

// Service is the interface definition for the user microservice
//
// CreateUser effectively add a user object to the backend data store
// it takes as input an type interface and returns the object id of the created
// user and any error encountered that may have occurred during this transaction
type Service interface {
	CreateUser(ctx context.Context, user interface{})(id string, err error)
}

// New returns a basic Service with all of the expected middlewares wired in.
func New(logger *zap.Logger, request, success, failed metrics.Counter) Service {
	var svc Service
	{
		svc = NewBasicService()
		svc = LoggingMiddleware(logger)(svc)
		svc = InstrumentingMiddleware( request, success, failed )(svc)
	}
	return svc
}

// NewBasicService returns a naïve, stateless implementation of Service.
func NewBasicService() Service {
	return basicService{}
}

type basicService struct{}

// CreateUser implements service.
//
// Creates a user in the backend store given some user object of interface type
func (basicService) CreateUser(ctx context.Context, user interface{}) (id string, err error) {
	if user == nil {
		return "", NullUser
	}
	// Todo: Implement logic for user service
	// panic("Implement me")
	return "hello", nil
}

