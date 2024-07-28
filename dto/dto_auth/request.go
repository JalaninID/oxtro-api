package dto_auth

import (
	authv1 "app/gen/proto/auth/v1"
	"app/model"
	"time"

	"connectrpc.com/connect"
	"github.com/google/uuid"
)

func RegisterToModel(req *connect.Request[authv1.RegisterRequest]) model.User {
	user := model.User{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		DeletedAt: nil,
		Name:      req.Msg.Name,
		Profile:   req.Msg.Profile,
		Username:  req.Msg.Username,
		Password:  req.Msg.Password,
		IsPremium: false,
		Gender:    req.Msg.Gender.String(),
	}
	user.StatusActive()
	return user
}
