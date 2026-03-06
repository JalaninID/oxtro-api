package main

import (
	"app/config"
	"app/plugin"
	"app/plugins/sample_crm"
	"app/routers"
	"context"
)

func main() {
	conf := config.NewConfig()

	// Initialize the plugin system
	hooks := plugin.NewHookEngine()
	store := plugin.NewGormPluginStore(conf.Database)
	manager := plugin.NewManager(conf.Database, conf.Logger, store, hooks)

	// Register plugins
	manager.Register(&sample_crm.CRMPlugin{})

	// Restore plugin states from DB and activate previously-active plugins
	if err := manager.LoadAndActivateAll(context.Background()); err != nil {
		conf.Logger.Warnf("plugin load error: %v", err)
	}

	router := routers.NewRouter(conf, hooks)
	router.RouterSetup()
	router.RouterAuth()
	router.RouterPlugin(manager)

	// Register routes from all active plugins
	manager.RegisterRoutes(router.Mux)

	router.Run()
}
