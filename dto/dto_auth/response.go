package dto_auth

import (
	authv1 "app/gen/proto/auth/v1"
	userv1 "app/gen/proto/user/v1"
	"app/model"

	"connectrpc.com/connect"
)

func RegisterResponse(user model.User) *connect.Response[authv1.RegisterResponse] {
	return &connect.Response[authv1.RegisterResponse]{
		Msg: &authv1.RegisterResponse{
			Id:        user.ID,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
			Username:  user.Username,
			Name:      user.Name,
			Profile:   user.Profile,
			Status:    user.Status,
			Premium:   user.IsPremium,
			Gender:    userv1.GENDER(userv1.GENDER_value[user.Gender]),
		},
	}
}
