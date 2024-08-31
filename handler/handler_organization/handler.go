package handler_organization

import (
	"app/domain"
	organizationv1 "app/gen/organization/v1"
	"app/gen/organization/v1/organizationv1connect"
	"context"

	"connectrpc.com/connect"
)

type Organization struct {
	serviceOrg domain.ServiceOrganization
	organizationv1connect.UnimplementedOrganizationHandler
}

func NewHandlerOrganization(serviceOrg domain.ServiceOrganization) *Organization {
	return &Organization{serviceOrg: serviceOrg}
}

func (h *Organization) CreateOrganization(ctx context.Context, req *connect.Request[organizationv1.RequestOrganization]) (*connect.Response[organizationv1.ResponseOrganization], error) {
	response, err := h.serviceOrg.CreateOrganization(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}
