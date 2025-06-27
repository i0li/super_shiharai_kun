package errors

import "errors"

var (
	ErrUnauthorized       = errors.New("unouthorized")
	ErrEmailAlreadyExists = errors.New("email already exist")
)
