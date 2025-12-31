package errors

import "errors"

var (
	ErrInternalServer  = errors.New("internal server error")
	ErrSQL             = errors.New("sql error")
	ErrTooManyRequests = errors.New("too many requests")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrInvalidToken    = errors.New("invalid token")
	ErrForbidden       = errors.New("forbidden")
	ErrNotFound        = errors.New("not found")
	ErrBadRequest      = errors.New("bad request")
)

var GeneralErrors = []error{
	ErrInternalServer,
	ErrSQL,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
	ErrNotFound,
	ErrBadRequest,
}
