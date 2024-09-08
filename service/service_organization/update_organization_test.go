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

func TestUpdateOrganization(t *testing.T) {
	ctx := context.Background()
	t.Run("success", func(t *testing.T) {
		mockRepoOrg := new(mocks.RepositoryOrganization)
		service := NewService(mockRepoOrg, &logrus.Logger{})
		mockRepoOrg.On("DetailOrganization", ctx, mock.AnythingOfType("model.FilterOrganization")).
			Return(model.Organization{UUID: "uuid"}, nil)
		mockRepoOrg.On("UpdateOrganization", ctx, mock.AnythingOfType("model.FilterOrganization"), mock.AnythingOfType("model.Organization")).
			Return(model.Organization{UUID: "uuid"}, nil)

		result, err := service.UpdateOrganization(ctx, &organizationv1.RequestOrganization{
			Id:   1,
			Uuid: "uuid",
		})
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Error Not Found Organization", func(t *testing.T) {
		mockRepoOrg := new(mocks.RepositoryOrganization)
		service := NewService(mockRepoOrg, &logrus.Logger{})
		mockRepoOrg.On("DetailOrganization", ctx, mock.AnythingOfType("model.FilterOrganization")).
			Return(model.Organization{}, gorm.ErrRecordNotFound)
		result, err := service.UpdateOrganization(ctx, &organizationv1.RequestOrganization{
			Id:   1,
			Uuid: "uuid",
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("Error Update Organization", func(t *testing.T) {
		mockRepoOrg := new(mocks.RepositoryOrganization)
		service := NewService(mockRepoOrg, &logrus.Logger{})
		mockRepoOrg.On("DetailOrganization", ctx, mock.AnythingOfType("model.FilterOrganization")).
			Return(model.Organization{UUID: "uuid"}, nil)
		mockRepoOrg.On("UpdateOrganization", ctx, mock.AnythingOfType("model.FilterOrganization"), mock.AnythingOfType("model.Organization")).
			Return(model.Organization{}, errors.New("error"))
		result, err := service.UpdateOrganization(ctx, &organizationv1.RequestOrganization{
			Id:   1,
			Uuid: "uuid",
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
