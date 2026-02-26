package constant

import "errors"

var (
	// ErrUserNotFound is a constant of error message when user not found status code : 404
	ErrUserNotFound         = errors.New("user not found")
	// ErrPassword is a constant of error message when password not match status code : 400
	ErrPassword             = errors.New("password not match")
	// ErrInvalidPaymentAmount is a constant of error message when invalid payment amount status code : 400
	ErrInvalidPaymentAmount = errors.New("invalid payment amount")
	// ErrNotPremium is a constant of error message when user not premium status code : 400
	ErrNotPremium           = errors.New("user not premium")
	// ErrInvalidCredentials is a constant of error message when credentials are invalid status code : 401
	ErrInvalidCredentials = errors.New("invalid credentials")
	// ErrUserLocked is a constant of error message when account temporarily locked status code : 423
	ErrUserLocked = errors.New("user account is temporarily locked")
	// ErrPasswordPolicy is a constant of error message when password does not satisfy policy status code : 400
	ErrPasswordPolicy = errors.New("password does not satisfy policy")
	// ErrEmailNotVerified is a constant of error message when email is not yet verified status code : 403
	ErrEmailNotVerified = errors.New("email is not verified")
)
