package service_organization

import (
	"app/constant"
	organizationv1 "app/gen/organization/v1"
	"app/model"
	"context"

	"connectrpc.com/connect"
	"gorm.io/gorm"
)

func (s *service) UpdateOrganization(ctx context.Context, req *organizationv1.RequestOrganization) (*organizationv1.ResponseOrganization, error) {
	orgDetail, err := s.repoOrg.DetailOrganization(ctx, model.FilterOrganization{
		ID:     int(req.Id),
		UUID:   req.Uuid,
		Domain: req.Domain,
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, connect.NewError(connect.CodeNotFound, constant.ErrOrganizationNotFound)
		}
		s.logger.Errorf("error get organization: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	org := FormatterUpdateRequestOrganization(req)
	org.Domain = orgDetail.Domain
	res, err := s.repoOrg.UpdateOrganization(ctx, model.FilterOrganization{
		ID: orgDetail.ID,
	}, org)
	if err != nil {
		s.logger.Errorf("error update organization: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}

	return FormatterResponseOrganization(res), nil
}
