package service_auth

import (
	"app/constant"
	authv1 "app/gen/proto/auth/v1"
	"app/pkg"
	"context"

	"connectrpc.com/connect"
	"gorm.io/gorm"
)

func (s *ServiceAuth) Login(ctx context.Context, req *connect.Request[authv1.LoginRequest]) (*connect.Response[authv1.LoginResponse], error) {
	user, err := s.repoUser.DetailByUsername(ctx, req.Msg.Username)
	if err == gorm.ErrRecordNotFound {
		return nil, connect.NewError(connect.CodeNotFound, constant.ErrUserNotFound)
	}
	if err != nil {
		s.logger.Error("s.repoUser.DetailByUsername", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	if err := pkg.ComparePassword(user.Password, req.Msg.Password); err != nil {
		return nil, connect.NewError(connect.CodeInternal, constant.ErrPassword)
	}

	token, err := pkg.Sign(map[string]any{
		"id": user.ID,
	}, 0)
	if err != nil {
		s.logger.Error("pkg.Sign", err)
		return nil, constant.ErrInternalServer
	}
	msg := &authv1.LoginResponse{
		Token: token,
	}
	return &connect.Response[authv1.LoginResponse]{
		Msg: msg,
	}, nil
}
