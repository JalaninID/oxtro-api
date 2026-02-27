package plugin

import (
	"context"
	"sort"
	"sync"
)

// ActionFunc is a callback for action events (fire-and-forget).
// Analogous to WordPress add_action callback.
type ActionFunc func(ctx context.Context, payload any) error

// FilterFunc is a callback that transforms data passing through.
// Analogous to WordPress add_filter callback.
// It receives data and must return the (possibly modified) data.
type FilterFunc func(ctx context.Context, data any) (any, error)

type hookEntry[T any] struct {
	pluginID string
	priority int
	fn       T
}

// HookEngine manages actions and filters, providing WordPress-like extensibility.
type HookEngine struct {
	mu      sync.RWMutex
	actions map[string][]hookEntry[ActionFunc]
	filters map[string][]hookEntry[FilterFunc]
}

// NewHookEngine creates a new hook engine.
func NewHookEngine() *HookEngine {
	return &HookEngine{
		actions: make(map[string][]hookEntry[ActionFunc]),
		filters: make(map[string][]hookEntry[FilterFunc]),
	}
}

// AddAction registers a callback for a named action.
// Priority determines execution order (lower = earlier, default 10).
// Analogous to WordPress add_action().
func (h *HookEngine) AddAction(hookName string, pluginID string, priority int, fn ActionFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.actions[hookName] = append(h.actions[hookName], hookEntry[ActionFunc]{
		pluginID: pluginID,
		priority: priority,
		fn:       fn,
	})
	sort.Slice(h.actions[hookName], func(i, j int) bool {
		return h.actions[hookName][i].priority < h.actions[hookName][j].priority
	})
}

// DoAction fires all callbacks registered for a named action.
// Errors from individual callbacks are collected but do not stop execution.
// Analogous to WordPress do_action().
func (h *HookEngine) DoAction(ctx context.Context, hookName string, payload any) []error {
	h.mu.RLock()
	entries := make([]hookEntry[ActionFunc], len(h.actions[hookName]))
	copy(entries, h.actions[hookName])
	h.mu.RUnlock()

	var errs []error
	for _, entry := range entries {
		if err := entry.fn(ctx, payload); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// AddFilter registers a filter callback that can transform data.
// Priority determines execution order (lower = earlier, default 10).
// Analogous to WordPress add_filter().
func (h *HookEngine) AddFilter(hookName string, pluginID string, priority int, fn FilterFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.filters[hookName] = append(h.filters[hookName], hookEntry[FilterFunc]{
		pluginID: pluginID,
		priority: priority,
		fn:       fn,
	})
	sort.Slice(h.filters[hookName], func(i, j int) bool {
		return h.filters[hookName][i].priority < h.filters[hookName][j].priority
	})
}

// ApplyFilters passes data through all registered filter callbacks in priority order.
// Each filter receives the output of the previous one.
// Analogous to WordPress apply_filters().
func (h *HookEngine) ApplyFilters(ctx context.Context, hookName string, data any) (any, error) {
	h.mu.RLock()
	entries := make([]hookEntry[FilterFunc], len(h.filters[hookName]))
	copy(entries, h.filters[hookName])
	h.mu.RUnlock()

	var err error
	for _, entry := range entries {
		data, err = entry.fn(ctx, data)
		if err != nil {
			return data, err
		}
	}
	return data, nil
}

// RemovePluginHooks removes all hooks (actions and filters) for a specific plugin.
// Called during plugin deactivation.
func (h *HookEngine) RemovePluginHooks(pluginID string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for name, entries := range h.actions {
		filtered := make([]hookEntry[ActionFunc], 0, len(entries))
		for _, e := range entries {
			if e.pluginID != pluginID {
				filtered = append(filtered, e)
			}
		}
		h.actions[name] = filtered
	}

	for name, entries := range h.filters {
		filtered := make([]hookEntry[FilterFunc], 0, len(entries))
		for _, e := range entries {
			if e.pluginID != pluginID {
				filtered = append(filtered, e)
			}
		}
		h.filters[name] = filtered
	}
}

// HasAction checks if any callbacks are registered for the given action hook.
func (h *HookEngine) HasAction(hookName string) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.actions[hookName]) > 0
}

// HasFilter checks if any callbacks are registered for the given filter hook.
func (h *HookEngine) HasFilter(hookName string) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.filters[hookName]) > 0
}
