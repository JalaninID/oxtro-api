package plugin

import (
	"context"
	"net/http"
)

// PluginState represents the current state of a plugin in its lifecycle.
type PluginState string

const (
	StateInstalled PluginState = "installed"
	StateActive    PluginState = "active"
	StateInactive  PluginState = "inactive"
	StateError     PluginState = "error"
)

// Manifest describes a plugin's metadata, analogous to WordPress plugin headers.
type Manifest struct {
	// Unique identifier, e.g. "com.oxtro.crm"
	ID string `json:"id" yaml:"id"`
	// Human-readable name
	Name string `json:"name" yaml:"name"`
	// Semantic version, e.g. "1.2.0"
	Version string `json:"version" yaml:"version"`
	// Plugin author
	Author string `json:"author" yaml:"author"`
	// Short description
	Description string `json:"description" yaml:"description"`
	// Minimum oxtro-api version required
	MinAppVersion string `json:"min_app_version" yaml:"min_app_version"`
	// IDs of plugins this depends on (must be active before this one activates)
	Dependencies []string `json:"dependencies" yaml:"dependencies"`
	// Plugin category: "business", "integration", "workflow", "utility"
	Category string `json:"category" yaml:"category"`
	// URL to plugin's homepage or repository
	Homepage string `json:"homepage" yaml:"homepage"`
	// License identifier (e.g., "MIT", "Apache-2.0")
	License string `json:"license" yaml:"license"`
	// Permissions this plugin requires (e.g., ["database:write", "hooks:auth"])
	Permissions []string `json:"permissions" yaml:"permissions"`
}

// Plugin is the primary interface every plugin must implement.
// Analogous to WordPress's plugin activation/deactivation hooks.
type Plugin interface {
	// Manifest returns the plugin's metadata.
	Manifest() Manifest

	// OnInstall is called when the plugin is first installed.
	// Use this for one-time setup like creating default config values.
	OnInstall(ctx context.Context, pctx *PluginContext) error

	// OnActivate is called each time the plugin is activated.
	// Use this to register hooks, start background tasks, etc.
	OnActivate(ctx context.Context, pctx *PluginContext) error

	// OnDeactivate is called when the plugin is deactivated.
	// Data should be preserved (user may reactivate later).
	OnDeactivate(ctx context.Context, pctx *PluginContext) error

	// OnUninstall is called when the plugin is permanently removed.
	// Use this to clean up all plugin data.
	OnUninstall(ctx context.Context, pctx *PluginContext) error
}

// RouteRegistrar is an optional interface plugins implement to register HTTP routes.
// Plugin routes are automatically prefixed with /plugins/{plugin_id}/.
type RouteRegistrar interface {
	RegisterRoutes(mux *http.ServeMux, pctx *PluginContext)
}

// MigrationProvider is an optional interface for plugins that need database migrations.
type MigrationProvider interface {
	// MigrationDir returns the path to the plugin's SQL migration files.
	MigrationDir() string
	// MigrationTablePrefix returns the prefix for all plugin tables (e.g., "plg_crm_").
	MigrationTablePrefix() string
}

// HookSubscriber is an optional interface for plugins that want to listen to system events.
type HookSubscriber interface {
	// SubscribeHooks is called during activation to register action and filter hooks.
	SubscribeHooks(hooks *HookEngine)
}
