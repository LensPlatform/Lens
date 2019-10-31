package service

import "context"

// NotificationService describes the service.
type NotificationService interface {
	// Methods associated with user onboarding & password/email reset
	SendEmail(ctx context.Context, email string, content string) (string, error)
	SendWelcomeToLensEmail(ctx context.Context, email string) (string, error)
	SendPasswordChangeEmail(ctx context.Context, email string) (string, error)
	SendEmailAccountResetEmail(ctx context.Context, oldEmail string, newEmail string) (string, error)
}

type basicNotificationService struct{}

func (b *basicNotificationService) SendEmail(ctx context.Context, email string, content string) (string, error) {
	// TODO implement the business logic of SendEmail
	return "", nil
}
func (b *basicNotificationService) SendWelcomeToLensEmail(ctx context.Context, email string) (string, error) {
	// TODO implement the business logic of SendWelcomeToLensEmail
	return "", nil
}
func (b *basicNotificationService) SendPasswordChangeEmail(ctx context.Context, email string) (string, error) {
	// TODO implement the business logic of SendPasswordChangeEmail
	return "", nil
}
func (b *basicNotificationService) SendEmailAccountResetEmail(ctx context.Context, oldEmail string, newEmail string) (string, error) {
	// TODO implement the business logic of SendEmailAccountResetEmail
	return "", nil
}

// NewBasicNotificationService returns a naive, stateless implementation of NotificationService.
func NewBasicNotificationService() NotificationService {
	return &basicNotificationService{}
}

// New returns a NotificationService with all of the expected middleware wired in.
func New(middleware []Middleware) NotificationService {
	var svc NotificationService = NewBasicNotificationService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
