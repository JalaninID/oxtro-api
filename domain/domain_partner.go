package domain

import (
	"app/model"
	"context"
)
//go:generate mockery --name RepositoryInterest
type RepositoryInterest interface {
	Create(ctx context.Context, interest model.Interest) (model.Interest, error)
	ListPartnerSwipe(filter model.FilterInterest) ([]string, error)
	DeletePartnerExpired() error
}
