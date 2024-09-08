package service_organization

import (
	"app/domain/mocks"
	organizationv1 "app/gen/organization/v1"
	"app/model"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListOrganization(t *testing.T) {
	ctx := context.Background()
	t.Run("Test List Organization Success", func(t *testing.T) {
		mockRepoOrg := &mocks.RepositoryOrganization{}
		mockServiceOrg := NewService(mockRepoOrg, &logrus.Logger{})
		expectedOrg := []model.Organization{
			{
				ID:          1,
				Name:        "Test Org",
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
				DeletedAt:   nil,
				UUID:        uuid.New().String(),
				Domain:      "test.com",
				Logo:        "test.com/logo.png",
				Description: "Test Description",
			},
			{
				ID:          2,
				Name:        "Test Org 2",
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
				DeletedAt:   nil,
				UUID:        uuid.New().String(),
				Domain:      "test2.com",
				Logo:        "test2.com/logo.png",
				Description: "Test Description 2",
			},
		}
		mockRepoOrg.On("ListOrganization", ctx, mock.Anything).Return(expectedOrg, len(expectedOrg), nil)
		result, err := mockServiceOrg.ListOrganization(ctx, &organizationv1.ParamsOrganization{})
		assert.NoError(t, err)
		assert.Equal(t, len(expectedOrg), len(result.Organizations))
		mockRepoOrg.AssertExpectations(t)
	})
	t.Run("Test List Organization Failed", func(t *testing.T) {
		mockRepoOrg := &mocks.RepositoryOrganization{}
		mockServiceOrg := NewService(mockRepoOrg, &logrus.Logger{})
		mockRepoOrg.On("ListOrganization", ctx, mock.Anything).Return(nil, 0, errors.New("error"))
		result, err := mockServiceOrg.ListOrganization(ctx, &organizationv1.ParamsOrganization{})
		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepoOrg.AssertExpectations(t)
	})
}
