package handler_plugin

import (
	pluginv1 "app/gen/plugin/v1"
	"app/gen/plugin/v1/pluginv1connect"
	toolsv1 "app/gen/tools/v1"
	"app/plugin"
	"context"

	"connectrpc.com/connect"
)

type PluginHandler struct {
	manager *plugin.Manager
	pluginv1connect.UnimplementedPluginManagerHandler
}

func NewHandlerPlugin(manager *plugin.Manager) *PluginHandler {
	return &PluginHandler{manager: manager}
}

func (h *PluginHandler) ListPlugins(_ context.Context, _ *connect.Request[pluginv1.ListPluginsRequest]) (*connect.Response[pluginv1.ListPluginsResponse], error) {
	list := h.manager.ListPlugins()

	var plugins []*pluginv1.PluginInfo
	for _, info := range list {
		plugins = append(plugins, &pluginv1.PluginInfo{
			Id:           info.Manifest.ID,
			Name:         info.Manifest.Name,
			Version:      info.Manifest.Version,
			Author:       info.Manifest.Author,
			Description:  info.Manifest.Description,
			Category:     info.Manifest.Category,
			State:        string(info.State),
			Homepage:     info.Manifest.Homepage,
			License:      info.Manifest.License,
			Dependencies: info.Manifest.Dependencies,
			Permissions:  info.Manifest.Permissions,
		})
	}

	return connect.NewResponse(&pluginv1.ListPluginsResponse{Plugins: plugins}), nil
}

func (h *PluginHandler) InstallPlugin(ctx context.Context, req *connect.Request[pluginv1.PluginActionRequest]) (*connect.Response[pluginv1.PluginActionResponse], error) {
	if err := h.manager.Install(ctx, req.Msg.GetPluginId()); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&pluginv1.PluginActionResponse{
		Success: true,
		Message: "plugin installed",
	}), nil
}

func (h *PluginHandler) ActivatePlugin(ctx context.Context, req *connect.Request[pluginv1.PluginActionRequest]) (*connect.Response[pluginv1.PluginActionResponse], error) {
	if err := h.manager.Activate(ctx, req.Msg.GetPluginId()); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&pluginv1.PluginActionResponse{
		Success: true,
		Message: "plugin activated",
	}), nil
}

func (h *PluginHandler) DeactivatePlugin(ctx context.Context, req *connect.Request[pluginv1.PluginActionRequest]) (*connect.Response[pluginv1.PluginActionResponse], error) {
	if err := h.manager.Deactivate(ctx, req.Msg.GetPluginId()); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&pluginv1.PluginActionResponse{
		Success: true,
		Message: "plugin deactivated",
	}), nil
}

func (h *PluginHandler) UninstallPlugin(ctx context.Context, req *connect.Request[pluginv1.PluginActionRequest]) (*connect.Response[toolsv1.Empty], error) {
	if err := h.manager.Uninstall(ctx, req.Msg.GetPluginId()); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&toolsv1.Empty{}), nil
}

func (h *PluginHandler) GetPluginConfig(_ context.Context, _ *connect.Request[pluginv1.GetPluginConfigRequest]) (*connect.Response[pluginv1.GetPluginConfigResponse], error) {
	// TODO: implement when plugin config store is wired
	return connect.NewResponse(&pluginv1.GetPluginConfigResponse{}), nil
}

func (h *PluginHandler) SetPluginConfig(_ context.Context, _ *connect.Request[pluginv1.SetPluginConfigRequest]) (*connect.Response[pluginv1.PluginActionResponse], error) {
	// TODO: implement when plugin config store is wired
	return connect.NewResponse(&pluginv1.PluginActionResponse{
		Success: true,
		Message: "config updated",
	}), nil
}
