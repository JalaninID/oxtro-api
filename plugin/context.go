package plugin

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// PluginContext provides controlled access to core services for plugins.
// This is passed to every plugin lifecycle method.
type PluginContext struct {
	// DB provides database access for the plugin.
	DB *gorm.DB

	// Logger is a plugin-scoped logger that automatically tags entries with the plugin ID.
	Logger *logrus.Entry

	// Hooks provides access to the hook engine for firing actions/filters.
	Hooks *HookEngine

	// Config provides per-plugin key-value configuration storage.
	Config PluginConfigStore

	// PluginID is the unique identifier of the current plugin.
	PluginID string
}

// PluginConfigStore provides key-value configuration storage per plugin.
type PluginConfigStore interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
	GetAll() (map[string]string, error)
}
