package service_organization

import (
	"app/constant"
	organizationv1 "app/gen/organization/v1"
	"app/model"
	"context"

	"connectrpc.com/connect"
	"gorm.io/gorm"
)

func (s *service) DetailOrganization(ctx context.Context, params *organizationv1.ParamsOrganization) (*organizationv1.ResponseOrganization, error) {
	org, err := s.serviceOrg.DetailOrganization(ctx, model.FilterOrganization{
		ID:     int(params.Id),
		UUID:   params.Uuid,
		Name:   params.Name,
		Domain: params.Domain,
	})
	if err == gorm.ErrRecordNotFound {
		return nil, connect.NewError(connect.CodeNotFound, constant.ErrOrganizationNotFound)
	}
	if err != nil {
		s.logger.Errorf("error get detail organization : %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	return FormatterResponseOrganization(org), nil
}
