package helper

import (
	"errors"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrNoUserProvided = errors.New("user not provided")
	ErrDBConnection = errors.New("database connection error")
	ErrPasswordsNotEqual = errors.New("password not equal to confirmed password")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrNoUsernameProvided = errors.New("no user name provided")
	ErrNoPasswordProvided = errors.New("no password provided")
	ErrInvalidUsernameProvided = errors.New("invalid username provided")
	ErrInvalidPasswordProvided = errors.New("invalid password provided")

	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)
