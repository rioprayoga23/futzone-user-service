package errors

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrUsernameExists       = errors.New("username exists")
	ErrPasswordIncorrect    = errors.New("password incorrect")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrUsernameExists,
	ErrPasswordIncorrect,
	ErrPasswordDoesNotMatch,
}
