package routers

import (
	"app/gen/proto/organization/v1/orgranizationv1connect"
	"app/injector"
)

func (r *Router) RouterOrganization() {
	orgranizationv1connect.NewOrganizationHandler(
		injector.InitializedOrganization(),
	)
}
