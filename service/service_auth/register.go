package service_auth

import (
	"app/constant"
	"app/dto/dto_auth"
	authv1 "app/gen/proto/auth/v1"
	"app/pkg"
	"context"

	"connectrpc.com/connect"
	"gorm.io/gorm"
)

func (s *ServiceAuth) Register(ctx context.Context, req *connect.Request[authv1.RegisterRequest]) (*connect.Response[authv1.RegisterResponse], error) {
	user, err := s.repoUser.DetailByUsername(ctx, req.Msg.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.logger.Error("s.repoUser.DetailByUsername", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	if user.ID != "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, constant.ErrUsernameExist)
	}
	userModel := dto_auth.RegisterToModel(req)
	password, err := pkg.HashPassword(userModel.Password)
	if err != nil {
		s.logger.Error("pkg.HashPassword", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	userModel.Password = password

	result, err := s.repoUser.Create(ctx, userModel)
	if err != nil {
		s.logger.Error("s.repoUser.Create", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	return dto_auth.RegisterResponse(result), nil
}
