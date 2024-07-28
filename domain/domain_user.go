package domain

import (
	"app/model"
	"context"
)

//go:generate mockery --name RepositoryUser
type RepositoryUser interface {
	DetailByID(ctx context.Context, id string) (model.User, error)
	Create(ctx context.Context, user model.User) (model.User, error)
	DetailByUsername(ctx context.Context, username string) (model.User, error)
	SearchPartner(filter model.FilterInterest) ([]model.User, int, error)
	UpdatePurchasePremium(ctx context.Context, userID string, premium bool) error
}
