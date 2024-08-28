package constant

import "errors"

var (
	// ErrAuthorization is a constant of error message when authorization failed status code : 401
	ErrAuthorization = errors.New("Authorization failed")
)
