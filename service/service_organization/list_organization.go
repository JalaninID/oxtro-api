package service_organization

import (
	"app/constant"
	organizationv1 "app/gen/organization/v1"
	toolsv1 "app/gen/tools/v1"
	"app/model"
	"app/pkg"
	"context"
	"math"

	"connectrpc.com/connect"
)

func (s *service) ListOrganization(ctx context.Context, req *organizationv1.ParamsOrganization) (*organizationv1.ResponseOrganizationList, error) {
	pagination := pkg.PaginationBuilder(int(req.PerPage), int(req.Page))
	organizations, total, err := s.repoOrg.ListOrganization(ctx, model.FilterOrganization{
		ID:      int(req.Id),
		Name:    req.Name,
		UUID:    req.Uuid,
		Domain:  req.Domain,
		Offset:  pagination.Offset,
		PerPage: pagination.PerPage,
	})
	if err != nil {
		s.logger.Errorf("error get list organization: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	totalPage := int(math.Ceil(float64(total) / float64(pagination.PerPage)))
	response := FormatterResponseOrganizationList(organizations)
	response.Pagination = &toolsv1.Pagination{
		Total:     int32(total),
		PerPage:   int32(pagination.PerPage),
		TotalPage: int32(totalPage),
		Page:      int32(pagination.Page),
	}

	return response, nil
}
