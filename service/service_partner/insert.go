package service_partner

import (
	"app/constant"
	"app/dto/dto_partner"
	partnerv1 "app/gen/proto/partner/v1"
	"app/pkg"
	"context"

	"connectrpc.com/connect"
)

func (s *ServicePartner) SwipePartner(ctx context.Context, req *connect.Request[partnerv1.RequestSwipePartner]) (*connect.Response[partnerv1.ResponseSwipePartner], error) {
	userMeta, ok := ctx.Value("user").(pkg.MetaToken)
	if !ok {
		return nil, connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
	}
	dtoReq := dto_partner.RequestSwipePartner(req)
	dtoReq.UserID = userMeta.ID
	result, err := s.repoPartner.Create(ctx, dtoReq)
	if err != nil {
		s.logger.Error("s.repoPartner.Create", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	return &connect.Response[partnerv1.ResponseSwipePartner]{
		Msg: &partnerv1.ResponseSwipePartner{
			PartnerId:  result.InterestUserID,
			IsInterest: result.IsInterest,
		},
	}, nil
}
