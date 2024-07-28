package service_partner

import (
	"app/constant"
	"app/domain/mocks"
	partnerv1 "app/gen/proto/partner/v1"
	"app/model"
	"app/pkg"
	"context"
	"math"
	"testing"

	"connectrpc.com/connect"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAllCaseSearchPartner(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"Success SearchPartner":  testServicePartnerSearchPartner_Success,
		"User Limit Not Premium": testServicePartnerSearchPartner_UserLimitNotPremium,
		"User Premium Not Limit": testServicePartnerSearchPartner_UserPremiumNotLimit,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}
func testServicePartnerSearchPartner_Success(t *testing.T) {
	mockRepoUser := new(mocks.RepositoryUser)
	mockRepoPartner := new(mocks.RepositoryInterest)
	svc := NewService(&logrus.Logger{}, mockRepoPartner, mockRepoUser)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "user", pkg.MetaToken{
		ID: "user_id",
	})

	req := &connect.Request[partnerv1.RequestSearchPartner]{
		Msg: &partnerv1.RequestSearchPartner{
			PerPage: 10,
			Page:    1,
		},
	}

	expectedUser := model.User{
		ID:     "user_id",
		Gender: "MALE",
	}

	expectedNotUser := []string{
		"partner_id_1",
	}

	expectedResults := []model.User{
		{
			ID:     "partner_id_1",
			Gender: "FEMALE",
		},
		{
			ID:     "partner_id_2",
			Gender: "FEMALE",
		},
	}

	mockRepoUser.On("DetailByID", ctx, "user_id").Return(expectedUser, nil)
	mockRepoPartner.On("ListPartnerSwipe", model.FilterInterest{UserID: "user_id"}).Return(expectedNotUser, nil)
	mockRepoUser.On("SearchPartner", mock.AnythingOfType("model.FilterInterest")).Return(expectedResults, len(expectedResults)-1, nil)

	res, err := svc.SearchPartner(ctx, req)
	assert.NotNil(t, res)
	assert.NoError(t, err)
	assert.Equal(t, len(expectedResults), len(res.Msg.Users))
	assert.Equal(t, int32(1), res.Msg.Pagination.Page)
	assert.Equal(t, int32(10), res.Msg.Pagination.PerPage)
	assert.Equal(t, int32(math.Ceil(float64(len(expectedResults))/float64(10))), res.Msg.Pagination.TotalPage)
	assert.Equal(t, int32(len(expectedResults)-1), res.Msg.Pagination.Total)

	mockRepoUser.AssertExpectations(t)
	mockRepoPartner.AssertExpectations(t)
}
func testServicePartnerSearchPartner_UserLimitNotPremium(t *testing.T) {
	mockRepoUser := new(mocks.RepositoryUser)
	mockRepoPartner := new(mocks.RepositoryInterest)
	svc := NewService(&logrus.Logger{}, mockRepoPartner, mockRepoUser)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "user", pkg.MetaToken{
		ID: "user_id",
	})

	req := &connect.Request[partnerv1.RequestSearchPartner]{
		Msg: &partnerv1.RequestSearchPartner{
			PerPage: 10,
			Page:    1,
		},
	}

	expectedUser := model.User{
		ID:        "user_id",
		Gender:    "MALE",
		IsPremium: false,
	}

	expectedNotUser := []string{
		"partner_id_1",
		"partner_id_2",
		"partner_id_3",
		"partner_id_4",
		"partner_id_5",
		"partner_id_6",
		"partner_id_7",
		"partner_id_8",
		"partner_id_9",
		"partner_id_10",
		"partner_id_11",
	}

	_ = []model.User{
		{
			ID:     "partner_id_1",
			Gender: "FEMALE",
		},
		{
			ID:     "partner_id_2",
			Gender: "FEMALE",
		},
		{
			ID:     "partner_id_3",
			Gender: "FEAMLE",
		},
		{
			ID:     "partner_id_5",
			Gender: "FEAMLE",
		},
		{
			ID:     "partner_id_6",
			Gender: "FEAMLE",
		},
		{
			ID:     "partner_id_7",
			Gender: "FEAMLE",
		},
		{
			ID:     "partner_id_8",
			Gender: "FEAMLE",
		},
		{
			ID:     "partner_id_9",
			Gender: "FEAMLE",
		},
		{
			ID:     "partner_id_10",
			Gender: "FEAMLE",
		},
		{
			ID:     "partner_id_11",
			Gender: "FEAMLE",
		},
		{
			ID:     "partner_id_12",
			Gender: "FEAMLE",
		},
	}

	mockRepoUser.On("DetailByID", ctx, "user_id").Return(expectedUser, nil)
	mockRepoPartner.On("ListPartnerSwipe", model.FilterInterest{UserID: "user_id"}).Return(expectedNotUser, nil)
	// mockRepoUser.On("SearchPartner", mock.AnythingOfType("model.FilterInterest")).Return(expectedResults, len(expectedResults)-1, nil)

	res, err := svc.SearchPartner(ctx, req)
	assert.Nil(t, res)
	assert.ErrorIs(t, err, constant.ErrNotPremium)
	assert.Error(t, err)
	mockRepoUser.AssertExpectations(t)
	mockRepoPartner.AssertExpectations(t)
}
func testServicePartnerSearchPartner_UserPremiumNotLimit(t *testing.T) {
	mockRepoUser := new(mocks.RepositoryUser)
	mockRepoPartner := new(mocks.RepositoryInterest)
	svc := NewService(&logrus.Logger{}, mockRepoPartner, mockRepoUser)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "user", pkg.MetaToken{
		ID: "user_id",
	})

	req := &connect.Request[partnerv1.RequestSearchPartner]{
		Msg: &partnerv1.RequestSearchPartner{
			PerPage: 10,
			Page:    1,
		},
	}

	expectedUser := model.User{
		ID:        "user_id",
		Gender:    "MALE",
		IsPremium: true,
	}

	expectedNotUser := []string{
		"partner_id_1",
	}

	expectedResults := []model.User{
		{
			ID:     "partner_id_1",
			Gender: "FEMALE",
		},
		{
			ID:     "partner_id_2",
			Gender: "FEMALE",
		},
		{
			ID:     "partner_id_3",
			Gender: "FEAMLE",
		},
		{
			ID:     "partner_id_5",
			Gender: "FEAMLE",
		},
		{
			ID:     "partner_id_6",
			Gender: "FEAMLE",
		},
		{
			ID:     "partner_id_7",
			Gender: "FEAMLE",
		},
		{
			ID:     "partner_id_8",
			Gender: "FEAMLE",
		},
		{
			ID:     "partner_id_9",
			Gender: "FEAMLE",
		},
		{
			ID:     "partner_id_10",
			Gender: "FEAMLE",
		},
		{
			ID:     "partner_id_11",
			Gender: "FEAMLE",
		},
		{
			ID:     "partner_id_12",
			Gender: "FEAMLE",
		},
	}

	mockRepoUser.On("DetailByID", ctx, "user_id").Return(expectedUser, nil)
	mockRepoPartner.On("ListPartnerSwipe", model.FilterInterest{UserID: "user_id"}).Return(expectedNotUser, nil)
	mockRepoUser.On("SearchPartner", mock.AnythingOfType("model.FilterInterest")).Return(expectedResults, len(expectedResults)-1, nil)

	res, err := svc.SearchPartner(ctx, req)
	assert.NotNil(t, res, "res is nil")
	assert.NoError(t, err, "err is not nil")
	assert.Equal(t, len(expectedResults), len(res.Msg.Users), "len(expectedResults) != len(res.Msg.Users)")
	mockRepoUser.AssertExpectations(t)
	mockRepoPartner.AssertExpectations(t)
}
