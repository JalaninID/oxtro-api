package mocks

import (
	"context"
)

// HookDispatcher is a no-op implementation for testing.
type HookDispatcher struct{}

func (m *HookDispatcher) DoAction(_ context.Context, _ string, _ any) []error {
	return nil
}

func (m *HookDispatcher) ApplyFilters(_ context.Context, _ string, data any) (any, error) {
	return data, nil
}
