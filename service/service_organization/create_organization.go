package service_organization

import (
	"app/constant"
	organizationv1 "app/gen/organization/v1"
	"app/model"
	"context"

	"connectrpc.com/connect"
	"gorm.io/gorm"
)

func (s *service) CreateOrganization(ctx context.Context, req *organizationv1.RequestOrganization) (*organizationv1.ResponseOrganization, error) {
	org := FormatterCreateRequestOrganization(req)

	orgDetail, err := s.serviceOrg.DetailOrganization(ctx, model.FilterOrganization{
		Domain: org.Domain,
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		s.logger.Errorf("error get organization: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	if orgDetail.ID != 0 {
		return nil, connect.NewError(connect.CodeAlreadyExists, constant.ErrOrganizationAlreadyExists)
	}

	result, err := s.serviceOrg.CreateOrganization(ctx, org)
	if err != nil {
		s.logger.Errorf("error create organization: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	return FormatterResponseOrganization(result), nil
}
