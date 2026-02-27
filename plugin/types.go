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

// UIManifest describes plugin-contributed UI metadata for host-rendered frontends.
// It intentionally stays schema-driven (no arbitrary JS) to keep UI secure and upgrade-safe.
type UIManifest struct {
	PluginID              string           `json:"plugin_id" yaml:"plugin_id"`
	UISchemaVersion       string           `json:"ui_schema_version" yaml:"ui_schema_version"`
	RequiresHostUIVersion string           `json:"requires_host_ui_version" yaml:"requires_host_ui_version"`
	Navigation            []NavigationItem `json:"navigation" yaml:"navigation"`
	Views                 []UIView         `json:"views" yaml:"views"`
	ViewExtensions        []ViewExtension  `json:"view_extensions" yaml:"view_extensions"`
	Actions               []UIAction       `json:"actions" yaml:"actions"`
}

type NavigationItem struct {
	ID                  string   `json:"id" yaml:"id"`
	Label               string   `json:"label" yaml:"label"`
	Icon                string   `json:"icon" yaml:"icon"`
	Path                string   `json:"path" yaml:"path"`
	Order               int32    `json:"order" yaml:"order"`
	RequiredPermissions []string `json:"required_permissions" yaml:"required_permissions"`
}

type UIView struct {
	ID         string     `json:"id" yaml:"id"`
	Type       string     `json:"type" yaml:"type"`
	RoutePath  string     `json:"route_path" yaml:"route_path"`
	Layout     string     `json:"layout" yaml:"layout"`
	Root       UINode     `json:"root" yaml:"root"`
	DataSource DataSource `json:"data_source" yaml:"data_source"`
}

type UINode struct {
	Component string            `json:"component" yaml:"component"`
	NodeID    string            `json:"node_id" yaml:"node_id"`
	Props     map[string]string `json:"props" yaml:"props"`
	Children  []UINode          `json:"children" yaml:"children"`
}

type DataSource struct {
	RPCMethod       string            `json:"rpc_method" yaml:"rpc_method"`
	RequestMapping  map[string]string `json:"request_mapping" yaml:"request_mapping"`
	ResponseMapping map[string]string `json:"response_mapping" yaml:"response_mapping"`
}

type ViewExtension struct {
	ID           string           `json:"id" yaml:"id"`
	TargetViewID string           `json:"target_view_id" yaml:"target_view_id"`
	Priority     int32            `json:"priority" yaml:"priority"`
	Operations   []PatchOperation `json:"operations" yaml:"operations"`
}

type PatchOperation struct {
	Op       PatchOp           `json:"op" yaml:"op"`
	Selector Selector          `json:"selector" yaml:"selector"`
	Node     UINode            `json:"node" yaml:"node"`
	SetProps map[string]string `json:"set_props" yaml:"set_props"`
}

type PatchOp string

const (
	PatchOpInsertBefore PatchOp = "insert_before"
	PatchOpInsertAfter  PatchOp = "insert_after"
	PatchOpReplace      PatchOp = "replace"
	PatchOpSetProps     PatchOp = "set_props"
	PatchOpRemove       PatchOp = "remove"
)

type Selector struct {
	By    SelectorBy `json:"by" yaml:"by"`
	Value string     `json:"value" yaml:"value"`
}

type SelectorBy string

const (
	SelectorByNodeID SelectorBy = "node_id"
	SelectorByPath   SelectorBy = "path"
)

type UIAction struct {
	ID                  string   `json:"id" yaml:"id"`
	Label               string   `json:"label" yaml:"label"`
	Type                string   `json:"type" yaml:"type"`
	RPCMethod           string   `json:"rpc_method" yaml:"rpc_method"`
	RequiredPermissions []string `json:"required_permissions" yaml:"required_permissions"`
	ConfirmMessage      string   `json:"confirm_message" yaml:"confirm_message"`
	SuccessToast        string   `json:"success_toast" yaml:"success_toast"`
}

// UIManifestProvider is an optional interface for plugins exposing host-rendered UI schema.
type UIManifestProvider interface {
	UIManifest() UIManifest
}
