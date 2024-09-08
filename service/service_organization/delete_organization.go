package service_organization

import (
	"app/constant"
	organizationv1 "app/gen/organization/v1"
	toolsv1 "app/gen/tools/v1"
	"app/model"
	"context"

	"connectrpc.com/connect"
	"gorm.io/gorm"
)

func (s *service) DeleteOrganization(ctx context.Context, req *organizationv1.ParamsOrganization) (*toolsv1.Empty, error) {
	org, err := s.repoOrg.DetailOrganization(ctx, model.FilterOrganization{
		UUID:   req.Uuid,
		ID:     int(req.Id),
		Domain: req.Domain,
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, connect.NewError(connect.CodeNotFound, constant.ErrOrganizationNotFound)
		}
		s.logger.Errorf("Error detail organization: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	if err := s.repoOrg.DeleteOrganization(ctx, model.FilterOrganization{ID: org.ID}); err != nil {
		s.logger.Errorf("Error delete organization: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	return &toolsv1.Empty{}, nil
}
