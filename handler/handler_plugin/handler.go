package handler_plugin

import (
	pluginv1 "app/gen/plugin/v1"
	"app/gen/plugin/v1/pluginv1connect"
	toolsv1 "app/gen/tools/v1"
	"app/plugin"
	"context"
	"errors"

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

func (h *PluginHandler) GetPluginUiManifest(_ context.Context, req *connect.Request[pluginv1.GetPluginUiManifestRequest]) (*connect.Response[pluginv1.GetPluginUiManifestResponse], error) {
	manifest, err := h.manager.GetPluginUIManifest(req.Msg.GetPluginId())
	if err != nil {
		if errors.Is(err, plugin.ErrPluginNotFound) || errors.Is(err, plugin.ErrPluginUIManifestNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pluginv1.GetPluginUiManifestResponse{
		Manifest: toProtoUIManifest(manifest),
	}), nil
}

func (h *PluginHandler) ListActivePluginUiManifests(_ context.Context, _ *connect.Request[pluginv1.ListActivePluginUiManifestsRequest]) (*connect.Response[pluginv1.ListActivePluginUiManifestsResponse], error) {
	list := h.manager.ListActivePluginUIManifests()
	manifests := make([]*pluginv1.PluginUiManifest, 0, len(list))
	for _, item := range list {
		manifests = append(manifests, toProtoUIManifest(item.Manifest))
	}

	return connect.NewResponse(&pluginv1.ListActivePluginUiManifestsResponse{
		Manifests: manifests,
	}), nil
}

func toProtoUIManifest(m plugin.UIManifest) *pluginv1.PluginUiManifest {
	return &pluginv1.PluginUiManifest{
		PluginId:              m.PluginID,
		UiSchemaVersion:       m.UISchemaVersion,
		RequiresHostUiVersion: m.RequiresHostUIVersion,
		Navigation:            toProtoNavigation(m.Navigation),
		Views:                 toProtoViews(m.Views),
		ViewExtensions:        toProtoViewExtensions(m.ViewExtensions),
		Actions:               toProtoUIActions(m.Actions),
	}
}

func toProtoNavigation(items []plugin.NavigationItem) []*pluginv1.NavigationItem {
	out := make([]*pluginv1.NavigationItem, 0, len(items))
	for _, item := range items {
		out = append(out, &pluginv1.NavigationItem{
			Id:                  item.ID,
			Label:               item.Label,
			Icon:                item.Icon,
			Path:                item.Path,
			Order:               item.Order,
			RequiredPermissions: item.RequiredPermissions,
		})
	}
	return out
}

func toProtoViews(views []plugin.UIView) []*pluginv1.UiView {
	out := make([]*pluginv1.UiView, 0, len(views))
	for _, v := range views {
		out = append(out, &pluginv1.UiView{
			Id:        v.ID,
			Type:      v.Type,
			RoutePath: v.RoutePath,
			Layout:    v.Layout,
			Root:      toProtoUINode(v.Root),
			DataSource: &pluginv1.DataSource{
				RpcMethod:       v.DataSource.RPCMethod,
				RequestMapping:  v.DataSource.RequestMapping,
				ResponseMapping: v.DataSource.ResponseMapping,
			},
		})
	}
	return out
}

func toProtoUINode(node plugin.UINode) *pluginv1.UiNode {
	children := make([]*pluginv1.UiNode, 0, len(node.Children))
	for _, child := range node.Children {
		children = append(children, toProtoUINode(child))
	}

	return &pluginv1.UiNode{
		Component: node.Component,
		NodeId:    node.NodeID,
		Props:     node.Props,
		Children:  children,
	}
}

func toProtoViewExtensions(extensions []plugin.ViewExtension) []*pluginv1.ViewExtension {
	out := make([]*pluginv1.ViewExtension, 0, len(extensions))
	for _, ext := range extensions {
		out = append(out, &pluginv1.ViewExtension{
			Id:           ext.ID,
			TargetViewId: ext.TargetViewID,
			Priority:     ext.Priority,
			Operations:   toProtoPatchOperations(ext.Operations),
		})
	}
	return out
}

func toProtoPatchOperations(operations []plugin.PatchOperation) []*pluginv1.PatchOperation {
	out := make([]*pluginv1.PatchOperation, 0, len(operations))
	for _, op := range operations {
		out = append(out, &pluginv1.PatchOperation{
			Op: toProtoPatchOp(op.Op),
			Selector: &pluginv1.Selector{
				By:    toProtoSelectorBy(op.Selector.By),
				Value: op.Selector.Value,
			},
			Node:     toProtoUINode(op.Node),
			SetProps: op.SetProps,
		})
	}
	return out
}

func toProtoPatchOp(op plugin.PatchOp) pluginv1.PatchOp {
	switch op {
	case plugin.PatchOpInsertBefore:
		return pluginv1.PatchOp_PATCH_OP_INSERT_BEFORE
	case plugin.PatchOpInsertAfter:
		return pluginv1.PatchOp_PATCH_OP_INSERT_AFTER
	case plugin.PatchOpReplace:
		return pluginv1.PatchOp_PATCH_OP_REPLACE
	case plugin.PatchOpSetProps:
		return pluginv1.PatchOp_PATCH_OP_SET_PROPS
	case plugin.PatchOpRemove:
		return pluginv1.PatchOp_PATCH_OP_REMOVE
	default:
		return pluginv1.PatchOp_PATCH_OP_UNSPECIFIED
	}
}

func toProtoSelectorBy(by plugin.SelectorBy) pluginv1.SelectorBy {
	switch by {
	case plugin.SelectorByNodeID:
		return pluginv1.SelectorBy_SELECTOR_BY_NODE_ID
	case plugin.SelectorByPath:
		return pluginv1.SelectorBy_SELECTOR_BY_PATH
	default:
		return pluginv1.SelectorBy_SELECTOR_BY_UNSPECIFIED
	}
}

func toProtoUIActions(actions []plugin.UIAction) []*pluginv1.UiAction {
	out := make([]*pluginv1.UiAction, 0, len(actions))
	for _, action := range actions {
		out = append(out, &pluginv1.UiAction{
			Id:                  action.ID,
			Label:               action.Label,
			Type:                action.Type,
			RpcMethod:           action.RPCMethod,
			RequiredPermissions: action.RequiredPermissions,
			ConfirmMessage:      action.ConfirmMessage,
			SuccessToast:        action.SuccessToast,
		})
	}
	return out
}
