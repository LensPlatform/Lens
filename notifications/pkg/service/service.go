package service

import (
	"golang.org/x/net/context"
)

// NotificationsService describes the service.
type NotificationsService interface {
	SendEmail(ctx context.Context, email string, content string) (string, error)
	SendWelcomeToLensEmail(ctx context.Context, email string) (string, error)
	SendPasswordChangeEmail(ctx context.Context, email string) (string, error)
	SendEmailAccountResetEmail(ctx context.Context, oldEmail string, newEmail string) (string, error)
}

type basicNotificationsService struct{}

func (b *basicNotificationsService) SendEmail(ctx context.Context, email string, content string) (s0 string, e1 error) {
	// TODO implement the business logic of SendEmail
	return s0, e1
}
func (b *basicNotificationsService) SendWelcomeToLensEmail(ctx context.Context, email string) (s0 string, e1 error) {
	// TODO implement the business logic of SendWelcomeToLensEmail
	return s0, e1
}
func (b *basicNotificationsService) SendPasswordChangeEmail(ctx context.Context, email string) (s0 string, e1 error) {
	// TODO implement the business logic of SendPasswordChangeEmail
	return s0, e1
}
func (b *basicNotificationsService) SendEmailAccountResetEmail(ctx context.Context, oldEmail string, newEmail string) (s0 string, e1 error) {
	// TODO implement the business logic of SendEmailAccountResetEmail
	return s0, e1
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
