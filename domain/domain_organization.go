package domain

import (
	organizationv1 "app/gen/organization/v1"
	"app/model"
	"context"
)

//go:generate mockery --name RepositoryOrganization
type RepositoryOrganization interface {
	CreateOrganization(ctx context.Context, organization model.Organization) (model.Organization, error)
	DetailOrganization(ctx context.Context, filter model.FilterOrganization) (model.Organization, error)
}

type ServiceOrganization interface {
	CreateOrganization(ctx context.Context, req *organizationv1.RequestOrganization) (*organizationv1.ResponseOrganization, error)
	DetailOrganization(ctx context.Context, params *organizationv1.ParamsOrganization) (*organizationv1.ResponseOrganization, error)
}
