package apperr

import "errors"

var (
	ErrUnauthorized       = errors.New("unauthorized")
	ErrEmailAlreadyExists = errors.New("email already exist")
)
