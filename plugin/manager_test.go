package plugin

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sirupsen/logrus"
)

// --- Mock Plugin ---

type mockPlugin struct {
	manifest        Manifest
	onInstallFn     func(ctx context.Context, pctx *PluginContext) error
	onActivateFn    func(ctx context.Context, pctx *PluginContext) error
	onDeactivateFn  func(ctx context.Context, pctx *PluginContext) error
	onUninstallFn   func(ctx context.Context, pctx *PluginContext) error
	subscribeHooksFn func(hooks *HookEngine)
}

func newMockPlugin(id string) *mockPlugin {
	return &mockPlugin{
		manifest: Manifest{
			ID:      id,
			Name:    id,
			Version: "1.0.0",
		},
	}
}

func (p *mockPlugin) Manifest() Manifest { return p.manifest }

func (p *mockPlugin) OnInstall(ctx context.Context, pctx *PluginContext) error {
	if p.onInstallFn != nil {
		return p.onInstallFn(ctx, pctx)
	}
	return nil
}

func (p *mockPlugin) OnActivate(ctx context.Context, pctx *PluginContext) error {
	if p.onActivateFn != nil {
		return p.onActivateFn(ctx, pctx)
	}
	return nil
}

func (p *mockPlugin) OnDeactivate(ctx context.Context, pctx *PluginContext) error {
	if p.onDeactivateFn != nil {
		return p.onDeactivateFn(ctx, pctx)
	}
	return nil
}

func (p *mockPlugin) OnUninstall(ctx context.Context, pctx *PluginContext) error {
	if p.onUninstallFn != nil {
		return p.onUninstallFn(ctx, pctx)
	}
	return nil
}

func (p *mockPlugin) SubscribeHooks(hooks *HookEngine) {
	if p.subscribeHooksFn != nil {
		p.subscribeHooksFn(hooks)
	}
}

// --- Mock PluginStore ---

type mockPluginStore struct {
	mu      sync.Mutex
	records map[string]PluginRecord
}

func newMockPluginStore() *mockPluginStore {
	return &mockPluginStore{
		records: make(map[string]PluginRecord),
	}
}

func (s *mockPluginStore) SavePluginRecord(ctx context.Context, record PluginRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.records[record.PluginID] = record
	return nil
}

func (s *mockPluginStore) UpdatePluginState(ctx context.Context, pluginID string, state PluginState) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if r, ok := s.records[pluginID]; ok {
		r.State = state
		s.records[pluginID] = r
	}
	return nil
}

func (s *mockPluginStore) DeletePluginRecord(ctx context.Context, pluginID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.records, pluginID)
	return nil
}

func (s *mockPluginStore) ListPluginRecords(ctx context.Context) ([]PluginRecord, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var records []PluginRecord
	for _, r := range s.records {
		records = append(records, r)
	}
	return records, nil
}

func (s *mockPluginStore) GetPluginRecord(ctx context.Context, pluginID string) (PluginRecord, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	r, ok := s.records[pluginID]
	if !ok {
		return PluginRecord{}, errors.New("not found")
	}
	return r, nil
}

// --- Helper ---

func newTestManager() (*Manager, *mockPluginStore) {
	logger := logrus.New()
	logger.SetLevel(logrus.ErrorLevel)
	store := newMockPluginStore()
	hooks := NewHookEngine()
	m := NewManager(nil, logger, store, hooks)
	return m, store
}

// --- Tests ---

func TestRegisterPlugin(t *testing.T) {
	m, _ := newTestManager()
	p := newMockPlugin("com.test.plugin")

	err := m.Register(p)
	require.NoError(t, err)

	// Duplicate registration should fail
	err = m.Register(p)
	assert.ErrorIs(t, err, ErrPluginAlreadyRegistered)
}

func TestInstallPlugin(t *testing.T) {
	m, store := newTestManager()
	ctx := context.Background()
	p := newMockPlugin("com.test.plugin")

	var installCalled bool
	p.onInstallFn = func(ctx context.Context, pctx *PluginContext) error {
		installCalled = true
		assert.Equal(t, "com.test.plugin", pctx.PluginID)
		return nil
	}

	err := m.Register(p)
	require.NoError(t, err)

	err = m.Install(ctx, "com.test.plugin")
	require.NoError(t, err)
	assert.True(t, installCalled)

	// Verify record in store
	record, err := store.GetPluginRecord(ctx, "com.test.plugin")
	require.NoError(t, err)
	assert.Equal(t, StateInstalled, record.State)
	assert.Equal(t, "1.0.0", record.Version)

	// Double install should fail
	err = m.Install(ctx, "com.test.plugin")
	assert.ErrorIs(t, err, ErrPluginAlreadyInstalled)
}

func TestInstallPluginFailure(t *testing.T) {
	m, _ := newTestManager()
	ctx := context.Background()
	p := newMockPlugin("com.test.failing")
	p.onInstallFn = func(ctx context.Context, pctx *PluginContext) error {
		return errors.New("install error")
	}

	_ = m.Register(p)
	err := m.Install(ctx, "com.test.failing")
	assert.ErrorIs(t, err, ErrPluginInstallFailed)

	state, _ := m.GetPluginState("com.test.failing")
	assert.Equal(t, StateError, state)
}

func TestActivatePlugin(t *testing.T) {
	m, store := newTestManager()
	ctx := context.Background()
	p := newMockPlugin("com.test.plugin")

	var activateCalled bool
	p.onActivateFn = func(ctx context.Context, pctx *PluginContext) error {
		activateCalled = true
		return nil
	}

	_ = m.Register(p)
	_ = m.Install(ctx, "com.test.plugin")

	err := m.Activate(ctx, "com.test.plugin")
	require.NoError(t, err)
	assert.True(t, activateCalled)

	state, _ := m.GetPluginState("com.test.plugin")
	assert.Equal(t, StateActive, state)

	// Verify store updated
	record, _ := store.GetPluginRecord(ctx, "com.test.plugin")
	assert.Equal(t, StateActive, record.State)

	// Double activation should fail
	err = m.Activate(ctx, "com.test.plugin")
	assert.ErrorIs(t, err, ErrPluginAlreadyActive)
}

func TestActivateWithDependency(t *testing.T) {
	m, _ := newTestManager()
	ctx := context.Background()

	base := newMockPlugin("com.test.base")
	dependent := newMockPlugin("com.test.dependent")
	dependent.manifest.Dependencies = []string{"com.test.base"}

	_ = m.Register(base)
	_ = m.Register(dependent)
	_ = m.Install(ctx, "com.test.base")
	_ = m.Install(ctx, "com.test.dependent")

	// Should fail: base not active
	err := m.Activate(ctx, "com.test.dependent")
	assert.ErrorIs(t, err, ErrPluginDependencyNotActive)

	// Activate base first
	_ = m.Activate(ctx, "com.test.base")

	// Now dependent should succeed
	err = m.Activate(ctx, "com.test.dependent")
	require.NoError(t, err)
}

func TestActivateRegistersHooks(t *testing.T) {
	m, _ := newTestManager()
	ctx := context.Background()
	p := newMockPlugin("com.test.hooks")

	p.subscribeHooksFn = func(hooks *HookEngine) {
		hooks.AddAction("test.custom", "com.test.hooks", 10, func(ctx context.Context, payload any) error {
			return nil
		})
	}

	_ = m.Register(p)
	_ = m.Install(ctx, "com.test.hooks")
	_ = m.Activate(ctx, "com.test.hooks")

	assert.True(t, m.Hooks().HasAction("test.custom"))
}

func TestDeactivatePlugin(t *testing.T) {
	m, _ := newTestManager()
	ctx := context.Background()
	p := newMockPlugin("com.test.plugin")

	var deactivateCalled bool
	p.onDeactivateFn = func(ctx context.Context, pctx *PluginContext) error {
		deactivateCalled = true
		return nil
	}
	p.subscribeHooksFn = func(hooks *HookEngine) {
		hooks.AddAction("test.action", "com.test.plugin", 10, func(ctx context.Context, payload any) error {
			return nil
		})
	}

	_ = m.Register(p)
	_ = m.Install(ctx, "com.test.plugin")
	_ = m.Activate(ctx, "com.test.plugin")

	err := m.Deactivate(ctx, "com.test.plugin")
	require.NoError(t, err)
	assert.True(t, deactivateCalled)

	state, _ := m.GetPluginState("com.test.plugin")
	assert.Equal(t, StateInactive, state)

	// Hooks should be removed
	assert.False(t, m.Hooks().HasAction("test.action"))
}

func TestDeactivateWithDependents(t *testing.T) {
	m, _ := newTestManager()
	ctx := context.Background()

	base := newMockPlugin("com.test.base")
	dependent := newMockPlugin("com.test.dependent")
	dependent.manifest.Dependencies = []string{"com.test.base"}

	_ = m.Register(base)
	_ = m.Register(dependent)
	_ = m.Install(ctx, "com.test.base")
	_ = m.Install(ctx, "com.test.dependent")
	_ = m.Activate(ctx, "com.test.base")
	_ = m.Activate(ctx, "com.test.dependent")

	// Should fail: dependent is still active
	err := m.Deactivate(ctx, "com.test.base")
	assert.ErrorIs(t, err, ErrPluginHasDependents)

	// Deactivate dependent first
	_ = m.Deactivate(ctx, "com.test.dependent")

	// Now base should succeed
	err = m.Deactivate(ctx, "com.test.base")
	require.NoError(t, err)
}

func TestUninstallPlugin(t *testing.T) {
	m, store := newTestManager()
	ctx := context.Background()
	p := newMockPlugin("com.test.plugin")

	var uninstallCalled bool
	p.onUninstallFn = func(ctx context.Context, pctx *PluginContext) error {
		uninstallCalled = true
		return nil
	}

	_ = m.Register(p)
	_ = m.Install(ctx, "com.test.plugin")
	_ = m.Activate(ctx, "com.test.plugin")

	// Uninstall should deactivate first, then uninstall
	err := m.Uninstall(ctx, "com.test.plugin")
	require.NoError(t, err)
	assert.True(t, uninstallCalled)

	// Plugin should be gone
	_, err = m.GetPluginState("com.test.plugin")
	assert.ErrorIs(t, err, ErrPluginNotFound)

	// Record should be gone from store
	_, err = store.GetPluginRecord(ctx, "com.test.plugin")
	assert.Error(t, err)
}

func TestInstallNotRegistered(t *testing.T) {
	m, _ := newTestManager()
	ctx := context.Background()

	err := m.Install(ctx, "com.test.nonexistent")
	assert.ErrorIs(t, err, ErrPluginNotFound)
}

func TestListPlugins(t *testing.T) {
	m, _ := newTestManager()
	ctx := context.Background()

	p1 := newMockPlugin("com.test.one")
	p2 := newMockPlugin("com.test.two")

	_ = m.Register(p1)
	_ = m.Register(p2)
	_ = m.Install(ctx, "com.test.one")
	_ = m.Install(ctx, "com.test.two")
	_ = m.Activate(ctx, "com.test.one")

	list := m.ListPlugins()
	assert.Len(t, list, 2)

	stateMap := make(map[string]PluginState)
	for _, info := range list {
		stateMap[info.Manifest.ID] = info.State
	}
	assert.Equal(t, StateActive, stateMap["com.test.one"])
	assert.Equal(t, StateInstalled, stateMap["com.test.two"])
}

func TestLoadAndActivateAll(t *testing.T) {
	m, store := newTestManager()
	ctx := context.Background()

	p := newMockPlugin("com.test.plugin")
	_ = m.Register(p)
	_ = m.Install(ctx, "com.test.plugin")

	// Simulate a previously active plugin in the store
	manifestJSON, _ := json.Marshal(p.Manifest())
	store.records["com.test.plugin"] = PluginRecord{
		PluginID:     "com.test.plugin",
		Version:      "1.0.0",
		State:        StateActive,
		ManifestJSON: manifestJSON,
	}

	// Reset state to simulate fresh startup
	m.setState("com.test.plugin", StateInstalled)

	err := m.LoadAndActivateAll(ctx)
	require.NoError(t, err)

	state, _ := m.GetPluginState("com.test.plugin")
	assert.Equal(t, StateActive, state)
}
