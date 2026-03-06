package handler_setup

import (
	"app/constant"
	"app/domain"
	"app/dto/dto_setup"
	setupv1 "app/gen/setup/v1"
	"app/gen/setup/v1/setupv1connect"
	toolsv1 "app/gen/tools/v1"
	"context"
	"errors"

	"connectrpc.com/connect"
)

type Setup struct {
	serviceSetup domain.ServiceSetup
	setupv1connect.UnimplementedSetupHandler
}

func NewHandlerSetup(serviceSetup domain.ServiceSetup) *Setup {
	return &Setup{serviceSetup: serviceSetup}
}

func (h *Setup) GetStatus(ctx context.Context, _ *connect.Request[toolsv1.Empty]) (*connect.Response[setupv1.SetupStatusResponse], error) {
	response, err := h.serviceSetup.GetStatus(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	return connect.NewResponse(&setupv1.SetupStatusResponse{
		IsCompleted: response.IsCompleted,
		DbHealthy:   response.DBHealthy,
		CompletedAt: response.CompletedAt,
	}), nil
}

func (h *Setup) RunSetup(ctx context.Context, req *connect.Request[setupv1.RunSetupRequest]) (*connect.Response[setupv1.RunSetupResponse], error) {
	response, err := h.serviceSetup.RunSetup(ctx, dto_setup.RunSetupRequest{
		Username: req.Msg.GetUsername(),
		Email:    req.Msg.GetEmail(),
		Password: req.Msg.GetPassword(),
	})
	if err != nil {
		switch {
		case errors.Is(err, constant.ErrInvalidCredentials), errors.Is(err, constant.ErrPasswordPolicy):
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		case errors.Is(err, constant.ErrSetupAlreadyCompleted):
			return nil, connect.NewError(connect.CodeFailedPrecondition, err)
		default:
			return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
		}
	}

	return connect.NewResponse(&setupv1.RunSetupResponse{
		Success: response.Success,
		Message: response.Message,
	}), nil
}
