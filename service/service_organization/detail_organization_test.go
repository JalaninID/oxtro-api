package service_organization

import (
	"app/constant"
	"app/domain/mocks"
	organizationv1 "app/gen/organization/v1"
	"app/model"
	"context"
	"errors"
	"testing"

	"connectrpc.com/connect"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestDetailOrganization(t *testing.T) {
	ctx := context.Background()
	params := &organizationv1.ParamsOrganization{
		Id:     1,
		Uuid:   "test-uuid",
		Name:   "Test Org",
		Domain: "test.com",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepoOrg := &mocks.RepositoryOrganization{}
		mockServiceOrg := NewService(mockRepoOrg, &logrus.Logger{})

		expectedOrg := model.Organization{ID: 1, Name: "Test Org"}
		mockRepoOrg.On("DetailOrganization", ctx, mock.Anything).Return(expectedOrg, nil)

		resp, err := mockServiceOrg.DetailOrganization(ctx, params)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		mockRepoOrg.AssertExpectations(t)
	})

	t.Run("Not Found", func(t *testing.T) {
		mockRepoOrg := &mocks.RepositoryOrganization{}
		mockServiceOrg := NewService(mockRepoOrg, &logrus.Logger{})

		mockRepoOrg.On("DetailOrganization", ctx, mock.Anything).Return(model.Organization{}, gorm.ErrRecordNotFound)

		resp, err := mockServiceOrg.DetailOrganization(ctx, params)

		assert.Error(t, err)
		assert.Nil(t, resp)
		var errorDetail *connect.Error
		errors.As(err, &errorDetail)

		assert.Contains(t, errorDetail.Message(), connect.NewError(connect.CodeNotFound, constant.ErrOrganizationNotFound).Message())
		mockRepoOrg.AssertExpectations(t)
	})

	t.Run("Error Internal", func(t *testing.T) {
		mockRepoOrg := &mocks.RepositoryOrganization{}
		mockServiceOrg := NewService(mockRepoOrg, &logrus.Logger{})

		mockRepoOrg.On("DetailOrganization", ctx, mock.Anything).Return(model.Organization{}, errors.New("internal error"))

		resp, err := mockServiceOrg.DetailOrganization(ctx, params)

		assert.Error(t, err)
		assert.Nil(t, resp)
		var errorDetail *connect.Error
		errors.As(err, &errorDetail)
		assert.Contains(t, errorDetail.Message(), connect.NewError(connect.CodeInternal, constant.ErrInternalServer).Message())
		mockRepoOrg.AssertExpectations(t)
	})
}
