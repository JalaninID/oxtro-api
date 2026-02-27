package routers

import (
	"app/gen/plugin/v1/pluginv1connect"
	"app/handler/handler_plugin"
	"app/middleware"
	"app/plugin"

	"connectrpc.com/connect"
)

func (r *Router) RouterPlugin(manager *plugin.Manager) {
	path, handler := pluginv1connect.NewPluginManagerHandler(
		handler_plugin.NewHandlerPlugin(manager),
		connect.WithInterceptors(middleware.NewAuthInterceptor(nil)),
	)
	r.Mux.Handle(path, handler)
}
