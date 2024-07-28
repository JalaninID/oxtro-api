package service_user

import (
	"app/constant"
	userv1 "app/gen/proto/user/v1"
	"app/pkg"
	"context"

	"connectrpc.com/connect"
)

func (s *ServiceUser) PurchasePremium(ctx context.Context, req *connect.Request[userv1.PurchaseRequest]) (*connect.Response[userv1.PurchaseResponse], error) {
	userMeta, ok := ctx.Value("user").(pkg.MetaToken)
	if !ok {
		return nil, connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
	}
	if req.Msg.PaymentAmount != 100000 {
		return nil, connect.NewError(connect.CodeInvalidArgument, constant.ErrInvalidPaymentAmount)
	}
	err := s.repo.UpdatePurchasePremium(ctx, userMeta.ID, true)
	if err != nil {
		s.logger.Error("s.repo.UpdatePurchasePremium", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	return &connect.Response[userv1.PurchaseResponse]{
		Msg: &userv1.PurchaseResponse{
			Success: true,
			Message: "Purchase premium success",
		},
	}, nil
}
