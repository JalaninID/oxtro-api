package constant

import "errors"

var (
	// ErrOrganizationNotFound is returned when the organization is not found
	ErrOrganizationNotFound = errors.New("organization not found")
	// ErrOrganizationAlreadyExists is returned when the organization already exists
	ErrOrganizationAlreadyExists = errors.New("organization already exists")
)
