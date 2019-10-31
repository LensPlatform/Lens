package service

import (
	"context"
)

// NotificationsService describes the service.
type NotificationsService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	SendEmail(ctx context.Context, email string, content string) (id string, err error)
}

type basicNotificationsService struct{}

func (b *basicNotificationsService) SendEmail(ctx context.Context, email string, content string) (id string, err error) {
	// TODO implement the business logic of SendEmail
	return id, err
}

// NewBasicNotificationsService returns a naive, stateless implementation of NotificationsService.
func NewBasicNotificationsService() NotificationsService {
	return &basicNotificationsService{}
}

// New returns a NotificationsService with all of the expected middleware wired in.
func New(middleware []Middleware) NotificationsService {
	var svc NotificationsService = NewBasicNotificationsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
