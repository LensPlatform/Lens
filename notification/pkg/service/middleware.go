package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(NotificationService) NotificationService

type loggingMiddleware struct {
	logger log.Logger
	next   NotificationService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a NotificationService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next NotificationService) NotificationService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) SendEmail(ctx context.Context, email string, content string) (string, error) {
	defer func() {
		l.logger.Log("method", "SendEmail", "email", email, "content", content)
	}()
	return l.next.SendEmail(ctx, email, content)
}
func (l loggingMiddleware) SendWelcomeToLensEmail(ctx context.Context, email string) (string, error) {
	defer func() {
		l.logger.Log("method", "SendWelcomeToLensEmail", "email", email)
	}()
	return l.next.SendWelcomeToLensEmail(ctx, email)
}
func (l loggingMiddleware) SendPasswordChangeEmail(ctx context.Context, email string) (string, error) {
	defer func() {
		l.logger.Log("method", "SendPasswordChangeEmail", "email", email)
	}()
	return l.next.SendPasswordChangeEmail(ctx, email)
}
func (l loggingMiddleware) SendEmailAccountResetEmail(ctx context.Context, oldEmail string, newEmail string) (string, error) {
	defer func() {
		l.logger.Log("method", "SendEmailAccountResetEmail", "oldEmail", oldEmail, "newEmail", newEmail)
	}()
	return l.next.SendEmailAccountResetEmail(ctx, oldEmail, newEmail)
}
