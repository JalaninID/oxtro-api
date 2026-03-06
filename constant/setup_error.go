package constant

import "errors"

var (
	ErrSetupAlreadyCompleted = errors.New("setup already completed")
	ErrSetupRequired         = errors.New("setup is required")
)
