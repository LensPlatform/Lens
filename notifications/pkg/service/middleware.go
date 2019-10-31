package service

import (
	log "github.com/go-kit/kit/log"
	context "golang.org/x/net/context"
)

type Middleware func(NotificationsService) NotificationsService

type loggingMiddleware struct {
	logger log.Logger
	next   NotificationsService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next NotificationsService) NotificationsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) SendEmail(ctx context.Context, email string, content string) (id string, err error) {
	defer func() {
		l.logger.Log("method", "SendEmail", "email", email, "content", content, "id", id, "err", err)
	}()
	return l.next.SendEmail(ctx, email, content)
}
