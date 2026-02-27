package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Manager orchestrates plugin lifecycle operations.
type Manager struct {
	mu       sync.RWMutex
	registry map[string]Plugin
	states   map[string]PluginState
	store    PluginStore
	hooks    *HookEngine
	db       *gorm.DB
	logger   *logrus.Logger
}

// NewManager creates a new plugin manager.
func NewManager(db *gorm.DB, logger *logrus.Logger, store PluginStore, hooks *HookEngine) *Manager {
	return &Manager{
		registry: make(map[string]Plugin),
		states:   make(map[string]PluginState),
		store:    store,
		hooks:    hooks,
		db:       db,
		logger:   logger,
	}
}

// Register makes a plugin known to the system.
// For in-process plugins, this is called at compile time in main.go.
func (m *Manager) Register(p Plugin) error {
	manifest := p.Manifest()
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.registry[manifest.ID]; exists {
		return fmt.Errorf("%w: %s", ErrPluginAlreadyRegistered, manifest.ID)
	}

	m.registry[manifest.ID] = p
	m.states[manifest.ID] = StateInstalled
	m.logger.WithField("plugin", manifest.ID).Info("plugin registered")
	return nil
}

// Install runs the one-time installation for a plugin.
func (m *Manager) Install(ctx context.Context, pluginID string) error {
	p, err := m.getPlugin(pluginID)
	if err != nil {
		return err
	}

	// Check if already installed in DB
	_, dbErr := m.store.GetPluginRecord(ctx, pluginID)
	if dbErr == nil {
		return fmt.Errorf("%w: %s", ErrPluginAlreadyInstalled, pluginID)
	}

	pctx := m.buildPluginContext(pluginID)

	// Run plugin migrations if it provides them
	if mp, ok := p.(MigrationProvider); ok {
		if err := RunPluginMigrations(m.db, mp); err != nil {
			m.setState(pluginID, StateError)
			return fmt.Errorf("%w: %s: %v", ErrPluginMigrationFailed, pluginID, err)
		}
	}

	if err := p.OnInstall(ctx, pctx); err != nil {
		m.setState(pluginID, StateError)
		return fmt.Errorf("%w: %s: %v", ErrPluginInstallFailed, pluginID, err)
	}

	// Marshal manifest to JSON for storage
	manifestJSON, _ := json.Marshal(p.Manifest())

	if err := m.store.SavePluginRecord(ctx, PluginRecord{
		PluginID:     pluginID,
		Version:      p.Manifest().Version,
		State:        StateInstalled,
		ManifestJSON: manifestJSON,
		ConfigJSON:   json.RawMessage("{}"),
	}); err != nil {
		return fmt.Errorf("saving plugin record: %w", err)
	}

	m.setState(pluginID, StateInstalled)
	m.logger.WithField("plugin", pluginID).Info("plugin installed")
	return nil
}

// Activate starts a plugin, calling OnActivate and registering its hooks.
func (m *Manager) Activate(ctx context.Context, pluginID string) error {
	p, err := m.getPlugin(pluginID)
	if err != nil {
		return err
	}

	if m.getState(pluginID) == StateActive {
		return fmt.Errorf("%w: %s", ErrPluginAlreadyActive, pluginID)
	}

	// Check dependencies are active
	for _, depID := range p.Manifest().Dependencies {
		if m.getState(depID) != StateActive {
			return fmt.Errorf("%w: %s requires %s", ErrPluginDependencyNotActive, pluginID, depID)
		}
	}

	pctx := m.buildPluginContext(pluginID)

	// Let plugin register its hooks
	if hs, ok := p.(HookSubscriber); ok {
		hs.SubscribeHooks(m.hooks)
	}

	if err := p.OnActivate(ctx, pctx); err != nil {
		m.setState(pluginID, StateError)
		// Remove hooks on failure
		m.hooks.RemovePluginHooks(pluginID)
		return fmt.Errorf("%w: %s: %v", ErrPluginActivateFailed, pluginID, err)
	}

	m.setState(pluginID, StateActive)
	_ = m.store.UpdatePluginState(ctx, pluginID, StateActive)
	m.hooks.DoAction(ctx, HookPluginActivated, p.Manifest())
	m.logger.WithField("plugin", pluginID).Info("plugin activated")
	return nil
}

// Deactivate stops a plugin without removing its data.
func (m *Manager) Deactivate(ctx context.Context, pluginID string) error {
	p, err := m.getPlugin(pluginID)
	if err != nil {
		return err
	}

	if m.getState(pluginID) != StateActive {
		return fmt.Errorf("%w: %s", ErrPluginNotActive, pluginID)
	}

	// Check no active plugin depends on this one
	m.mu.RLock()
	for id, state := range m.states {
		if state != StateActive || id == pluginID {
			continue
		}
		dep := m.registry[id]
		for _, d := range dep.Manifest().Dependencies {
			if d == pluginID {
				m.mu.RUnlock()
				return fmt.Errorf("%w: %s is required by %s", ErrPluginHasDependents, pluginID, id)
			}
		}
	}
	m.mu.RUnlock()

	pctx := m.buildPluginContext(pluginID)

	if err := p.OnDeactivate(ctx, pctx); err != nil {
		m.logger.WithField("plugin", pluginID).WithError(err).Error("plugin deactivation error")
	}

	// Remove all hooks registered by this plugin
	m.hooks.RemovePluginHooks(pluginID)

	m.setState(pluginID, StateInactive)
	_ = m.store.UpdatePluginState(ctx, pluginID, StateInactive)
	m.hooks.DoAction(ctx, HookPluginDeactivated, p.Manifest())
	m.logger.WithField("plugin", pluginID).Info("plugin deactivated")
	return nil
}

// Uninstall permanently removes a plugin and all its data.
func (m *Manager) Uninstall(ctx context.Context, pluginID string) error {
	// Must be deactivated first
	if m.getState(pluginID) == StateActive {
		if err := m.Deactivate(ctx, pluginID); err != nil {
			return err
		}
	}

	p, err := m.getPlugin(pluginID)
	if err != nil {
		return err
	}

	pctx := m.buildPluginContext(pluginID)

	if err := p.OnUninstall(ctx, pctx); err != nil {
		m.logger.WithField("plugin", pluginID).WithError(err).Error("plugin uninstall error")
	}

	// Roll back migrations if plugin provides them
	if mp, ok := p.(MigrationProvider); ok {
		if err := RollbackPluginMigrations(m.db, mp); err != nil {
			m.logger.WithField("plugin", pluginID).WithError(err).Warn("migration rollback failed")
		}
	}

	_ = m.store.DeletePluginRecord(ctx, pluginID)

	m.mu.Lock()
	delete(m.registry, pluginID)
	delete(m.states, pluginID)
	m.mu.Unlock()

	m.logger.WithField("plugin", pluginID).Info("plugin uninstalled")
	return nil
}

// LoadAndActivateAll restores plugin states from the database on server startup.
// Analogous to WordPress's plugins_loaded hook.
func (m *Manager) LoadAndActivateAll(ctx context.Context) error {
	records, err := m.store.ListPluginRecords(ctx)
	if err != nil {
		return fmt.Errorf("loading plugin records: %w", err)
	}

	for _, record := range records {
		if record.State == StateActive {
			if err := m.Activate(ctx, record.PluginID); err != nil {
				m.logger.WithField("plugin", record.PluginID).WithError(err).
					Error("failed to activate plugin on startup")
			}
		}
	}

	m.hooks.DoAction(ctx, HookPluginsLoaded, nil)
	return nil
}

// RegisterRoutes registers HTTP routes for all active plugins that implement RouteRegistrar.
func (m *Manager) RegisterRoutes(mux *http.ServeMux) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for id, p := range m.registry {
		if m.states[id] != StateActive {
			continue
		}
		rr, ok := p.(RouteRegistrar)
		if !ok {
			continue
		}

		pctx := m.buildPluginContext(id)
		pluginMux := http.NewServeMux()
		rr.RegisterRoutes(pluginMux, pctx)
		prefix := fmt.Sprintf("/plugins/%s/", id)
		mux.Handle(prefix, http.StripPrefix(prefix, pluginMux))
		m.logger.WithField("plugin", id).Infof("registered routes at %s", prefix)
	}
}

// ListPlugins returns information about all registered plugins.
type PluginInfo struct {
	Manifest Manifest    `json:"manifest"`
	State    PluginState `json:"state"`
}

func (m *Manager) ListPlugins() []PluginInfo {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var list []PluginInfo
	for id, p := range m.registry {
		list = append(list, PluginInfo{
			Manifest: p.Manifest(),
			State:    m.states[id],
		})
	}
	return list
}

// GetPluginState returns the current state of a plugin.
func (m *Manager) GetPluginState(pluginID string) (PluginState, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	state, ok := m.states[pluginID]
	if !ok {
		return "", ErrPluginNotFound
	}
	return state, nil
}

// Hooks returns the hook engine.
func (m *Manager) Hooks() *HookEngine {
	return m.hooks
}

func (m *Manager) getPlugin(id string) (Plugin, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	p, ok := m.registry[id]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrPluginNotFound, id)
	}
	return p, nil
}

func (m *Manager) getState(id string) PluginState {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.states[id]
}

func (m *Manager) setState(id string, state PluginState) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.states[id] = state
}

func (m *Manager) buildPluginContext(pluginID string) *PluginContext {
	return &PluginContext{
		DB:       m.db,
		Logger:   m.logger.WithField("plugin", pluginID),
		Hooks:    m.hooks,
		Config:   NewGormPluginConfigStore(m.db, pluginID),
		PluginID: pluginID,
	}
}
