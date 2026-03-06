package sample_crm

import (
	"app/model"
	"app/plugin"
	"app/plugins/sample_crm/gen/sample_crm/v1/sample_crmv1connect"
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
		Permissions:   []string{"database:write", "hooks:auth"},
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

}

// MigrationDir returns the path to CRM migration files.
func (p *CRMPlugin) MigrationDir() string {
	return "plugins/sample_crm/migrations"
}

// MigrationTablePrefix returns the prefix for CRM tables.
func (p *CRMPlugin) MigrationTablePrefix() string {
	return "plg_crm_"
}

// UIManifest exposes schema-driven UI metadata for Oxtro UI hosts.
func (p *CRMPlugin) UIManifest() plugin.UIManifest {
	return plugin.UIManifest{
		PluginID:              PluginID,
		UISchemaVersion:       "1.0.0",
		RequiresHostUIVersion: ">=1.0 <2.0",
		Navigation: []plugin.NavigationItem{
			{
				ID:                  "crm.contacts",
				Label:               "CRM Contacts",
				Icon:                "users",
				Path:                "/app/plugins/crm/contacts",
				Order:               40,
				RequiredPermissions: []string{"crm.contacts.read"},
			},
		},
		Views: []plugin.UIView{
			{
				ID:        "crm.contacts.list",
				Type:      "page",
				RoutePath: "/app/plugins/crm/contacts",
				Layout:    "default",
				Root: plugin.UINode{
					Component: "table",
					NodeID:    "crm_contacts_table",
					Props: map[string]string{
						"title":             "Contacts",
						"columns":           "[{\"key\":\"name\",\"label\":\"Name\"},{\"key\":\"email\",\"label\":\"Email\"},{\"key\":\"phone\",\"label\":\"Phone\"},{\"key\":\"company\",\"label\":\"Company\"}]",
						"searchEnabled":     "true",
						"paginationEnabled": "true",
					},
				},
				DataSource: plugin.DataSource{
					RPCMethod: "sample_crm.v1.CRM/ListContacts",
					RequestMapping: map[string]string{
						"page":     "$query.page",
						"per_page": "$query.per_page",
						"search":   "$query.search",
					},
					ResponseMapping: map[string]string{
						"rows":  "$.contacts",
						"total": "$.total",
					},
				},
			},
			{
				ID:        "crm.contacts.create",
				Type:      "page",
				RoutePath: "/app/plugins/crm/contacts/new",
				Layout:    "default",
				Root: plugin.UINode{
					Component: "form",
					NodeID:    "crm_contact_form_create",
					Props: map[string]string{
						"fields": "[{\"key\":\"name\",\"type\":\"text\",\"required\":true},{\"key\":\"email\",\"type\":\"email\",\"required\":true},{\"key\":\"phone\",\"type\":\"text\"},{\"key\":\"company\",\"type\":\"text\"},{\"key\":\"notes\",\"type\":\"textarea\"}]",
					},
				},
			},
		},
		Actions: []plugin.UIAction{
			{
				ID:                  "crm.createContact",
				Label:               "Save Contact",
				Type:                "rpc",
				RPCMethod:           "sample_crm.v1.CRM/CreateContact",
				RequiredPermissions: []string{"crm.contacts.write"},
				SuccessToast:        "Contact created",
			},
			{
				ID:                  "crm.deleteContact",
				Label:               "Delete Contact",
				Type:                "rpc",
				RPCMethod:           "sample_crm.v1.CRM/DeleteContact",
				RequiredPermissions: []string{"crm.contacts.delete"},
				ConfirmMessage:      "Are you sure to delete this contact?",
				SuccessToast:        "Contact deleted",
			},
		},
	}
}
