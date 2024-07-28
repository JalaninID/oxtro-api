package service_user

import (
	"app/constant"
	"app/domain/mocks"
	userv1 "app/gen/proto/user/v1"
	"app/pkg"
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestAllCaseUser(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"Success Purchase Premium Package": testUserService_Success,
		"Invalid Payment Amount":           testUserService_InvalidPaymentAmount,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testUserService_Success(t *testing.T) {
	mockRepo := new(mocks.RepositoryUser)
	srv := NewService(&logrus.Logger{}, mockRepo)

	ctx := context.WithValue(context.Background(), "user", pkg.MetaToken{ID: "user_id"})

	req := &connect.Request[userv1.PurchaseRequest]{
		Msg: &userv1.PurchaseRequest{
			PaymentAmount: 100000,
		},
	}

	mockRepo.On("UpdatePurchasePremium", ctx, "user_id", true).Return(nil).Once()

	expectedResponse := &connect.Response[userv1.PurchaseResponse]{
		Msg: &userv1.PurchaseResponse{
			Success: true,
			Message: "Purchase premium success",
		},
	}

	response, err := srv.PurchasePremium(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.Msg.Message, response.Msg.Message)
	assert.Equal(t, expectedResponse.Msg.Success, response.Msg.Success)

	mockRepo.AssertExpectations(t)
}
func testUserService_InvalidPaymentAmount(t *testing.T) {
	mockRepo := new(mocks.RepositoryUser)
	srv := NewService(&logrus.Logger{}, mockRepo)

	ctx := context.WithValue(context.Background(), "user", pkg.MetaToken{ID: "user_id"})

	req := &connect.Request[userv1.PurchaseRequest]{
		Msg: &userv1.PurchaseRequest{
			PaymentAmount: 10000,
		},
	}

	response, err := srv.PurchasePremium(ctx, req)
	assert.ErrorIs(t, err, constant.ErrInvalidPaymentAmount)
	assert.Error(t, err)
	assert.Nil(t, response)
	mockRepo.AssertExpectations(t)
}
