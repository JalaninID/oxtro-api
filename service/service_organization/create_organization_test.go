package service_organization

import (
	"app/domain/mocks"
	organizationv1 "app/gen/organization/v1"
	"app/model"
	"context"
	"errors"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateOrganization(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepoOrg := &mocks.RepositoryOrganization{}
		mockServiceOrg := NewService(mockRepoOrg, &logrus.Logger{})

		ctx := context.Background()
		req := &organizationv1.RequestOrganization{
			Domain: "test.com",
		}

		mockRepoOrg.On("DetailOrganization", ctx, mock.AnythingOfType("model.FilterOrganization")).
			Return(model.Organization{}, gorm.ErrRecordNotFound)
		mockRepoOrg.On("CreateOrganization", ctx, mock.AnythingOfType("model.Organization")).
			Return(model.Organization{Domain: "test.com"}, nil)

		resp, err := mockServiceOrg.CreateOrganization(ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, resp.Domain, req.Domain)
		mockRepoOrg.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		mockRepoOrg := &mocks.RepositoryOrganization{}
		mockServiceOrg := NewService(mockRepoOrg, &logrus.Logger{})

		ctx := context.Background()
		req := &organizationv1.RequestOrganization{
			Domain: "test.com",
		}

		mockRepoOrg.On("DetailOrganization", ctx, mock.AnythingOfType("model.FilterOrganization")).
			Return(model.Organization{}, gorm.ErrRecordNotFound)
		mockRepoOrg.On("CreateOrganization", ctx, mock.AnythingOfType("model.Organization")).
			Return(model.Organization{}, errors.New("create error"))

		resp, err := mockServiceOrg.CreateOrganization(ctx, req)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
		mockRepoOrg.AssertExpectations(t)
	})

}
