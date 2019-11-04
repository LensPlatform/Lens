package service

import (
	"errors"
	)

var (
	ErrNotFound        = errors.New("not found")
	ErrNoUserProvided = errors.New("user not provided")
	)
