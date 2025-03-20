package errors

import "errors"

var (
	ErrNotFound       = errors.New("resource not found")
	ErrInvalidRequest = errors.New("invalid request")
)