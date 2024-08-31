package service_organization

import (
	"app/constant"
	organizationv1 "app/gen/organization/v1"
	"context"

	"connectrpc.com/connect"
)

func (s *service) CreateOrganization(ctx context.Context, req *organizationv1.RequestOrganization) (*organizationv1.ResponseOrganization, error) {
	org := FormatterCreateRequestOrganization(req)
	result, err := s.serviceOrg.CreateOrganization(ctx, org)
	if err != nil {
		s.logger.Errorf("error create organization: %v", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	return FormatterResponseOrganization(result), nil
}
