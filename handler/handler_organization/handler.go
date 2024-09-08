package handler_organization

import (
	"app/domain"
	organizationv1 "app/gen/organization/v1"
	"app/gen/organization/v1/organizationv1connect"
	toolsv1 "app/gen/tools/v1"
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

func (h *Organization) DetailOrganization(ctx context.Context, req *connect.Request[organizationv1.ParamsOrganization]) (*connect.Response[organizationv1.ResponseOrganization], error) {
	response, err := h.serviceOrg.DetailOrganization(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (h *Organization) ListOrganization(ctx context.Context, req *connect.Request[organizationv1.ParamsOrganization]) (*connect.Response[organizationv1.ResponseOrganizationList], error) {
	response, err := h.serviceOrg.ListOrganization(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (h *Organization) UpdateOrganization(ctx context.Context, req *connect.Request[organizationv1.RequestOrganization]) (*connect.Response[organizationv1.ResponseOrganization], error) {
	response, err := h.serviceOrg.UpdateOrganization(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (h *Organization) DeleteOrganization(ctx context.Context, req *connect.Request[organizationv1.ParamsOrganization]) (*connect.Response[toolsv1.Empty], error) {
	response, err := h.serviceOrg.DeleteOrganization(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}
