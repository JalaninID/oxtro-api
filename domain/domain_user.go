package domain

import (
	"app/model"
	"context"
)

//go:generate mockery --name RepositoryUser
type RepositoryUser interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
}
