package plugin

// Core action hook names fired by the oxtro core system.
// Plugins subscribe to these via HookSubscriber.SubscribeHooks().
const (
	// Auth action hooks
	HookUserRegistered    = "user.registered"
	HookUserLoggedIn      = "user.logged_in"
	HookUserLoggedOut     = "user.logged_out"
	HookPasswordChanged   = "user.password_changed"
	HookEmailVerified     = "user.email_verified"

	// Plugin system action hooks
	HookPluginsLoaded     = "plugins.loaded"
	HookPluginActivated   = "plugin.activated"
	HookPluginDeactivated = "plugin.deactivated"
)

// Core filter hook names. Data passes through these and can be modified by plugins.
const (
	FilterRegisterData  = "user.register.data"
	FilterLoginResponse = "user.login.response"
)
