package service_partner

import (
	"app/constant"
	"app/dto/dto_partner"
	partnerv1 "app/gen/proto/partner/v1"
	userv1 "app/gen/proto/user/v1"
	utilv1 "app/gen/proto/utils/v1"
	"app/model"
	"app/pkg"
	"context"
	"math"

	"connectrpc.com/connect"
)

func (s *ServicePartner) SearchPartner(ctx context.Context, req *connect.Request[partnerv1.RequestSearchPartner]) (*connect.Response[partnerv1.ResponseSearchPartner], error) {
	userMeta, ok := ctx.Value("user").(pkg.MetaToken)
	if !ok {
		return nil, connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
	}

	user, err := s.repoUser.DetailByID(ctx, userMeta.ID)
	if err != nil {
		s.logger.Error("s.repoUser.DetailByID", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	gender := s.compareGender(userv1.GENDER_value[user.Gender])

	notUser, err := s.repoPartner.ListPartnerSwipe(model.FilterInterest{
		UserID: user.ID,
	})
	notUser = append(notUser, user.ID)
	if err != nil {
		s.logger.Error("s.repoPartner.ListPartnerSwipe", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	if !user.IsPremium && len(notUser) >= 10 {
		return nil, connect.NewError(connect.CodeInvalidArgument, constant.ErrNotPremium)
	}
	pagination := pkg.PaginationBuilder(int(req.Msg.PerPage), int(req.Msg.Page))
	filter := model.FilterInterest{
		UserID:         user.ID,
		PerPage:        pagination.PerPage,
		Offset:         pagination.Offset,
		Page:           pagination.Page,
		Gender:         gender,
		InterestUserID: notUser,
	}

	results, total, err := s.repoUser.SearchPartner(filter)
	if err != nil {
		s.logger.Error("s.repoPartner.SearchPartner", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	partners := dto_partner.ResponseParnerts(results)
	paginate := &utilv1.Pagination{
		Page:      int32(pagination.Page),
		PerPage:   int32(pagination.PerPage),
		TotalPage: int32(math.Ceil(float64(total) / float64(pagination.PerPage))),
		Total:     int32(total),
	}

	return &connect.Response[partnerv1.ResponseSearchPartner]{
		Msg: &partnerv1.ResponseSearchPartner{
			Users:      partners,
			Pagination: paginate,
		},
	}, nil
}

func (s *ServicePartner) compareGender(gender int32) string {
	switch gender {
	case 0:
		return userv1.GENDER_name[int32(1)]
	case 1:
		return userv1.GENDER_name[int32(0)]
	}
	return ""
}
