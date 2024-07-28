package constant

import "errors"

var (
	// ErrInternalServer is a constant of error message when internal server error
	ErrInternalServer = errors.New("internal server error")
	ErrUsernameExist  = errors.New("username already exist")
)
