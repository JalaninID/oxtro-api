package service_auth

import (
	"app/constant"
	"app/domain/mocks"
	authv1 "app/gen/proto/auth/v1"
	"app/model"
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAllCaseLogin(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"Success Login":         testServiceAuthLogin_Success,
		"Wrong Password":        testServiceAuthLogin_WrongPassword,
		"User Not Found":        testServiceAuthLogin_UserNotFound,
		"Internal Server Error": testServiceAuthLogin_InternalServerError,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}
func testServiceAuthLogin_Success(t *testing.T) {
	var mockRepoUser = new(mocks.RepositoryUser)
	var srv = NewService(&logrus.Logger{}, mockRepoUser)
	ctx := context.Background()
	username := "testuser"
	password := "1"
	id := uuid.NewString()

	mockRepoUser.On("DetailByUsername", ctx, username).Return(model.User{
		ID:       id,
		Username: username,
		Password: "$2a$04$.j8bw6aqScH/CeVXetnDQ.5s/6NZZubXxKrNAbthotJ2e7jl.8BXO",
	}, nil)

	req := &connect.Request[authv1.LoginRequest]{
		Msg: &authv1.LoginRequest{
			Username: username,
			Password: password,
		},
	}
	response, err := srv.Login(ctx, req)
	assert.NotEqual(t, "", response.Msg.Token)
	assert.NoError(t, err)
	mockRepoUser.AssertExpectations(t)
}
func testServiceAuthLogin_WrongPassword(t *testing.T) {
	var mockRepoUser = new(mocks.RepositoryUser)
	var srv = NewService(&logrus.Logger{}, mockRepoUser)
	ctx := context.Background()
	username := "testuser"
	password := "2"
	id := uuid.NewString()

	mockRepoUser.On("DetailByUsername", ctx, username).Return(model.User{
		ID:       id,
		Username: username,
		Password: "$2a$04$.j8bw6aqScH/CeVXetnDQ.5s/6NZZubXxKrNAbthotJ2e7jl.8BXO",
	}, nil)

	req := &connect.Request[authv1.LoginRequest]{
		Msg: &authv1.LoginRequest{
			Username: username,
			Password: password,
		},
	}
	response, err := srv.Login(ctx, req)
	assert.Error(t, err)
	assert.ErrorIs(t, err, constant.ErrPassword)
	assert.Nil(t, response)
	mockRepoUser.AssertExpectations(t)
}
func testServiceAuthLogin_UserNotFound(t *testing.T) {
	var mockRepoUser = new(mocks.RepositoryUser)
	var srv = NewService(&logrus.Logger{}, mockRepoUser)
	ctx := context.Background()
	username := "testuser"
	password := "1"

	mockRepoUser.On("DetailByUsername", ctx, username).Return(model.User{}, gorm.ErrRecordNotFound)

	req := &connect.Request[authv1.LoginRequest]{
		Msg: &authv1.LoginRequest{
			Username: username,
			Password: password,
		},
	}

	response, err := srv.Login(ctx, req)
	assert.Error(t, err)
	assert.ErrorIs(t, err, constant.ErrUserNotFound)
	assert.Nil(t, response)
	mockRepoUser.AssertExpectations(t)
}
func testServiceAuthLogin_InternalServerError(t *testing.T) {
	var mockRepoUser = new(mocks.RepositoryUser)
	var srv = NewService(&logrus.Logger{}, mockRepoUser)
	ctx := context.Background()
	username := "testuser"
	password := "1"

	mockRepoUser.On("DetailByUsername", ctx, username).Return(model.User{}, gorm.ErrInvalidDB)

	req := &connect.Request[authv1.LoginRequest]{
		Msg: &authv1.LoginRequest{
			Username: username,
			Password: password,
		},
	}

	response, err := srv.Login(ctx, req)
	assert.Error(t, err)
	assert.ErrorIs(t, err, constant.ErrInternalServer)
	assert.Nil(t, response)
	mockRepoUser.AssertExpectations(t)
}
