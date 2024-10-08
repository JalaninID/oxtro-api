package domain

import (
	organizationv1 "app/gen/organization/v1"
	toolsv1 "app/gen/tools/v1"
	"app/model"
	"context"
)

//go:generate mockery --name RepositoryOrganization
type RepositoryOrganization interface {
	CreateOrganization(ctx context.Context, organization model.Organization) (model.Organization, error)
	DetailOrganization(ctx context.Context, filter model.FilterOrganization) (model.Organization, error)
	ListOrganization(ctx context.Context, filter model.FilterOrganization) ([]model.Organization, int, error)
	DeleteOrganization(ctx context.Context, filter model.FilterOrganization) error
	UpdateOrganization(ctx context.Context, filter model.FilterOrganization, organization model.Organization) (model.Organization, error)
}

//go:generate mockery --name ServiceOrganization
type ServiceOrganization interface {
	CreateOrganization(ctx context.Context, req *organizationv1.RequestOrganization) (*organizationv1.ResponseOrganization, error)
	DetailOrganization(ctx context.Context, params *organizationv1.ParamsOrganization) (*organizationv1.ResponseOrganization, error)
	ListOrganization(ctx context.Context, req *organizationv1.ParamsOrganization) (*organizationv1.ResponseOrganizationList, error)
	UpdateOrganization(ctx context.Context, req *organizationv1.RequestOrganization) (*organizationv1.ResponseOrganization, error)
	DeleteOrganization(ctx context.Context, req *organizationv1.ParamsOrganization) (*toolsv1.Empty, error)
}
