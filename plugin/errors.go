package plugin

import "errors"

var (
	ErrPluginNotFound           = errors.New("plugin not found")
	ErrPluginAlreadyRegistered  = errors.New("plugin is already registered")
	ErrPluginAlreadyInstalled   = errors.New("plugin is already installed")
	ErrPluginAlreadyActive      = errors.New("plugin is already active")
	ErrPluginNotActive          = errors.New("plugin is not active")
	ErrPluginDependencyNotActive = errors.New("required plugin dependency is not active")
	ErrPluginHasDependents      = errors.New("cannot deactivate: other active plugins depend on this plugin")
	ErrPluginInstallFailed      = errors.New("plugin installation failed")
	ErrPluginActivateFailed     = errors.New("plugin activation failed")
	ErrPluginMigrationFailed    = errors.New("plugin migration failed")
)
