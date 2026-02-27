package domain

import "context"

// HookDispatcher allows services to fire hooks without depending on the plugin package.
type HookDispatcher interface {
	DoAction(ctx context.Context, hookName string, payload any) []error
	ApplyFilters(ctx context.Context, hookName string, data any) (any, error)
}
