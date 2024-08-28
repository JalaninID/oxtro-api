package constant

import "errors"

var (
	// ErrInternalServer is a constant of error message when internal server error status code : 500
	ErrInternalServer = errors.New("internal server error")
	// ErrUsernameExist is a constant of error message when username already exist status code : 400
	ErrUsernameExist = errors.New("username already exist")
)
