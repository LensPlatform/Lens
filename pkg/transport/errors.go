package transport

import (
	"context"
	"errors"

	"go.uber.org/zap"
)

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

type Handler struct {
	logger *zap.Logger
}

func NewTransportHandler(logger *zap.Logger) Handler{
	return Handler{
		logger: logger,
	}
}

func (handler Handler) Handle(ctx context.Context, err error) {
	handler.logger.Error("error", zap.Error(err))
}

// The ErrorHandlerFunc type is an adapter to allow the use of
// ordinary function as ErrorHandler. If f is a function
// with the appropriate signature, ErrorHandlerFunc(f) is a
// ErrorHandler that calls f.
type ErrorHandlerFunc func(ctx context.Context, err error)

// Handle calls f(ctx, err).
func (f ErrorHandlerFunc) Handle(ctx context.Context, err error) {
	f(ctx, err)
}