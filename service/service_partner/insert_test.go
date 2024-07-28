package service_partner

import (
	"app/domain/mocks"
	partnerv1 "app/gen/proto/partner/v1"
	"app/model"
	"app/pkg"
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAllCaseSwipePartner(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"Success SwipePartner": testServicePartnerSwipePartner_Success,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}
func testServicePartnerSwipePartner_Success(t *testing.T) {
	mockRepoPartner := new(mocks.RepositoryInterest)
	mockRepoUser := new(mocks.RepositoryUser)
	svc := NewService(&logrus.Logger{}, mockRepoPartner, mockRepoUser)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "user", pkg.MetaToken{
		ID: "user_id",
	})

	req := &connect.Request[partnerv1.RequestSwipePartner]{
		Msg: &partnerv1.RequestSwipePartner{
			PartnerId:  "partner_id",
			IsInterest: true,
		},
	}

	expectedRes := model.Interest{
		InterestUserID: "partner_id",
		IsInterest:     true,
	}

	mockRepoPartner.On("Create", ctx, mock.AnythingOfType("model.Interest")).Return(expectedRes, nil)

	res, err := svc.SwipePartner(ctx, req)
	assert.NotNil(t, res)
	assert.NoError(t, err)
	assert.Equal(t, expectedRes.InterestUserID, res.Msg.PartnerId)
	assert.Equal(t, expectedRes.IsInterest, res.Msg.IsInterest)
	mockRepoUser.AssertExpectations(t)
}
