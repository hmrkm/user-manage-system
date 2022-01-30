package domain

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrNotAuthorized = errors.New("not authorized")
	ErrForbidden     = errors.New("forbidden")
	ErrNotReaching   = errors.New("not reaching")
)
