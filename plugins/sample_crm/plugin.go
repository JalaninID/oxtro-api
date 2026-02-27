package sample_crm

import (
	"app/gen/sample_crm/v1/sample_crmv1connect"
	"app/model"
	"app/plugin"
	"context"
	"net/http"

	"github.com/google/uuid"
)

const PluginID = "com.oxtro.crm"

// CRMPlugin implements the Plugin interface for CRM functionality.
type CRMPlugin struct{}

func (p *CRMPlugin) Manifest() plugin.Manifest {
	return plugin.Manifest{
		ID:            PluginID,
		Name:          "Oxtro CRM",
		Version:       "1.0.0",
		Author:        "Oxtro Team",
		Description:   "Customer relationship management plugin for Oxtro",
		MinAppVersion: "0.1.0",
		Category:      "business",
		Homepage:      "https://github.com/oxtro/oxtro-crm",
		License:       "MIT",
		Permissions:   []string{"database:write", "hooks:auth", "hooks:organization"},
	}
}

func (p *CRMPlugin) OnInstall(_ context.Context, pctx *plugin.PluginContext) error {
	pctx.Logger.Info("CRM plugin installed")
	return nil
}

func (p *CRMPlugin) OnActivate(_ context.Context, pctx *plugin.PluginContext) error {
	pctx.Logger.Info("CRM plugin activated")
	return nil
}

func (p *CRMPlugin) OnDeactivate(_ context.Context, pctx *plugin.PluginContext) error {
	pctx.Logger.Info("CRM plugin deactivated")
	return nil
}

func (p *CRMPlugin) OnUninstall(_ context.Context, pctx *plugin.PluginContext) error {
	pctx.Logger.Info("CRM plugin uninstalled")
	return nil
}

// RegisterRoutes registers CRM API endpoints.
// Routes are automatically prefixed with /plugins/com.oxtro.crm/
func (p *CRMPlugin) RegisterRoutes(mux *http.ServeMux, pctx *plugin.PluginContext) {
	repo := newRepository(pctx.DB)
	h := newHandler(repo)
	path, httpHandler := sample_crmv1connect.NewCRMHandler(h)
	mux.Handle(path, httpHandler)
}

// SubscribeHooks registers event listeners for system events.
func (p *CRMPlugin) SubscribeHooks(hooks *plugin.HookEngine) {
	// Auto-create a CRM contact when a new user registers
	hooks.AddAction(plugin.HookUserRegistered, PluginID, 10, func(ctx context.Context, payload any) error {
		user, ok := payload.(model.User)
		if !ok {
			return nil
		}

		// We need DB access through the context, but for simplicity
		// we'll just log here. In a real implementation, you'd inject the repo.
		_ = CRMContact{
			UUID:    uuid.NewString(),
			Name:    user.Name,
			Email:   user.Email,
			Company: "",
			Notes:   "Auto-created from user registration",
		}
		// In production: repo.Create(ctx, contact)
		return nil
	})

	// React when an organization is created
	hooks.AddAction(plugin.HookOrgCreated, PluginID, 10, func(ctx context.Context, payload any) error {
		// Could link org data to CRM
		return nil
	})
}

// MigrationDir returns the path to CRM migration files.
func (p *CRMPlugin) MigrationDir() string {
	return "plugins/sample_crm/migrations"
}

// MigrationTablePrefix returns the prefix for CRM tables.
func (p *CRMPlugin) MigrationTablePrefix() string {
	return "plg_crm_"
}
