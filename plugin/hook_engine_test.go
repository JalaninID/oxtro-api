package plugin

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddActionAndDoAction(t *testing.T) {
	engine := NewHookEngine()
	ctx := context.Background()

	var called bool
	engine.AddAction("test.hook", "plugin-a", 10, func(ctx context.Context, payload any) error {
		called = true
		assert.Equal(t, "hello", payload)
		return nil
	})

	errs := engine.DoAction(ctx, "test.hook", "hello")
	assert.Empty(t, errs)
	assert.True(t, called)
}

func TestDoActionPriorityOrder(t *testing.T) {
	engine := NewHookEngine()
	ctx := context.Background()

	var order []string

	engine.AddAction("test.order", "plugin-b", 20, func(ctx context.Context, payload any) error {
		order = append(order, "second")
		return nil
	})
	engine.AddAction("test.order", "plugin-a", 5, func(ctx context.Context, payload any) error {
		order = append(order, "first")
		return nil
	})
	engine.AddAction("test.order", "plugin-c", 50, func(ctx context.Context, payload any) error {
		order = append(order, "third")
		return nil
	})

	engine.DoAction(ctx, "test.order", nil)
	assert.Equal(t, []string{"first", "second", "third"}, order)
}

func TestDoActionErrorsDoNotStopExecution(t *testing.T) {
	engine := NewHookEngine()
	ctx := context.Background()

	var callCount int

	engine.AddAction("test.error", "plugin-a", 10, func(ctx context.Context, payload any) error {
		callCount++
		return errors.New("fail")
	})
	engine.AddAction("test.error", "plugin-b", 20, func(ctx context.Context, payload any) error {
		callCount++
		return nil
	})

	errs := engine.DoAction(ctx, "test.error", nil)
	assert.Len(t, errs, 1)
	assert.Equal(t, 2, callCount)
}

func TestDoActionNoHandlers(t *testing.T) {
	engine := NewHookEngine()
	ctx := context.Background()

	errs := engine.DoAction(ctx, "nonexistent.hook", nil)
	assert.Empty(t, errs)
}

func TestAddFilterAndApplyFilters(t *testing.T) {
	engine := NewHookEngine()
	ctx := context.Background()

	engine.AddFilter("test.filter", "plugin-a", 10, func(ctx context.Context, data any) (any, error) {
		num := data.(int)
		return num * 2, nil
	})
	engine.AddFilter("test.filter", "plugin-b", 20, func(ctx context.Context, data any) (any, error) {
		num := data.(int)
		return num + 3, nil
	})

	result, err := engine.ApplyFilters(ctx, "test.filter", 5)
	require.NoError(t, err)
	assert.Equal(t, 13, result) // (5 * 2) + 3
}

func TestApplyFiltersStopsOnError(t *testing.T) {
	engine := NewHookEngine()
	ctx := context.Background()

	engine.AddFilter("test.filter.err", "plugin-a", 10, func(ctx context.Context, data any) (any, error) {
		return nil, errors.New("filter failed")
	})
	engine.AddFilter("test.filter.err", "plugin-b", 20, func(ctx context.Context, data any) (any, error) {
		t.Fatal("should not reach here")
		return data, nil
	})

	_, err := engine.ApplyFilters(ctx, "test.filter.err", "data")
	assert.Error(t, err)
	assert.Equal(t, "filter failed", err.Error())
}

func TestApplyFiltersNoHandlers(t *testing.T) {
	engine := NewHookEngine()
	ctx := context.Background()

	result, err := engine.ApplyFilters(ctx, "nonexistent.filter", "original")
	require.NoError(t, err)
	assert.Equal(t, "original", result)
}

func TestRemovePluginHooks(t *testing.T) {
	engine := NewHookEngine()
	ctx := context.Background()

	var calledA, calledB bool

	engine.AddAction("test.remove", "plugin-a", 10, func(ctx context.Context, payload any) error {
		calledA = true
		return nil
	})
	engine.AddAction("test.remove", "plugin-b", 10, func(ctx context.Context, payload any) error {
		calledB = true
		return nil
	})
	engine.AddFilter("test.remove.filter", "plugin-a", 10, func(ctx context.Context, data any) (any, error) {
		return "modified-by-a", nil
	})

	// Remove plugin-a hooks
	engine.RemovePluginHooks("plugin-a")

	engine.DoAction(ctx, "test.remove", nil)
	assert.False(t, calledA, "plugin-a action should not be called after removal")
	assert.True(t, calledB, "plugin-b action should still be called")

	result, err := engine.ApplyFilters(ctx, "test.remove.filter", "original")
	require.NoError(t, err)
	assert.Equal(t, "original", result, "plugin-a filter should not be applied after removal")
}

func TestHasActionAndHasFilter(t *testing.T) {
	engine := NewHookEngine()

	assert.False(t, engine.HasAction("test.has"))
	assert.False(t, engine.HasFilter("test.has"))

	engine.AddAction("test.has", "plugin-a", 10, func(ctx context.Context, payload any) error {
		return nil
	})
	engine.AddFilter("test.has.filter", "plugin-a", 10, func(ctx context.Context, data any) (any, error) {
		return data, nil
	})

	assert.True(t, engine.HasAction("test.has"))
	assert.True(t, engine.HasFilter("test.has.filter"))
}

func TestHookEngineConcurrency(t *testing.T) {
	engine := NewHookEngine()
	ctx := context.Background()

	var wg sync.WaitGroup
	for i := range 100 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			engine.AddAction("concurrent.hook", "plugin-"+string(rune('a'+n%26)), n, func(ctx context.Context, payload any) error {
				return nil
			})
			engine.DoAction(ctx, "concurrent.hook", n)
		}(i)
	}
	wg.Wait()
}
