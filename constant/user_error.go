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
)
