package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(NotificationsService) NotificationsService

type loggingMiddleware struct {
	logger log.Logger
	next   NotificationsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a NotificationsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next NotificationsService) NotificationsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) SendEmail(ctx context.Context, email string, content string) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "SendEmail", "email", email, "content", content, "s0", s0, "e1", e1)
	}()
	return l.next.SendEmail(ctx, email, content)
}
func (l loggingMiddleware) SendWelcomeToLensEmail(ctx context.Context, email string) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "SendWelcomeToLensEmail", "email", email, "s0", s0, "e1", e1)
	}()
	return l.next.SendWelcomeToLensEmail(ctx, email)
}
func (l loggingMiddleware) SendPasswordChangeEmail(ctx context.Context, email string) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "SendPasswordChangeEmail", "email", email, "s0", s0, "e1", e1)
	}()
	return l.next.SendPasswordChangeEmail(ctx, email)
}
func (l loggingMiddleware) SendEmailAccountResetEmail(ctx context.Context, oldEmail string, newEmail string) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "SendEmailAccountResetEmail", "oldEmail", oldEmail, "newEmail", newEmail, "s0", s0, "e1", e1)
	}()
	return l.next.SendEmailAccountResetEmail(ctx, oldEmail, newEmail)
}
