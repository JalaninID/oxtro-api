package service_organization

import (
	"app/gen/proto/organization/v1/orgranizationv1connect"
)

type Organization struct {
	orgranizationv1connect.UnimplementedOrganizationHandler
}

func NewOrganizationService() *Organization {
	return &Organization{}
}
