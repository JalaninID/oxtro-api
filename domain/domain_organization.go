package domain

import (
	"app/model"
	"context"
)
//go:generate mockery --name RepositoryOrganization
type RepositoryOrganization interface {
	CreateOrganization(ctx context.Context, organization model.Organization) (model.Organization, error)
}
