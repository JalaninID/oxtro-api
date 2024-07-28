package service_auth

import (
	"app/constant"
	"app/domain/mocks"
	authv1 "app/gen/proto/auth/v1"
	userv1 "app/gen/proto/user/v1"
	"app/model"
	"context"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAllCaseRegister(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"Success Register":      testServiceAuthRegister_Success,
		"User Already Exist":    testServiceAuthRegister_UserAlreadyExist,
		"Internal Server Error": testServiceAuthRegister_InternalServerError,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}
func testServiceAuthRegister_Success(t *testing.T) {
	var mockRepoUser = new(mocks.RepositoryUser)
	var srv = NewService(&logrus.Logger{}, mockRepoUser)
	ctx := context.Background()

	req := authv1.RegisterRequest{
		Username: "testuser",
		Password: "testpassword",
		Gender:   userv1.GENDER_FEMALE,
		Name:     "testname",
		Profile:  "testprofile",
		Premium:  false,
	}
	connectReq := &connect.Request[authv1.RegisterRequest]{
		Msg: &req,
	}

	newUser := model.User{
		Username:  req.Username,
		Name:      req.Name,
		Profile:   req.Profile,
		Status:    "active",
		IsPremium: req.Premium,
		Gender:    userv1.GENDER_name[int32(req.Gender)],
	}
	mockRepoUser.On("DetailByUsername", ctx, connectReq.Msg.Username).Return(model.User{}, nil).Once()
	mockRepoUser.On("Create", ctx, mock.AnythingOfType("model.User")).Return(newUser, nil).Once()
	response, err := srv.Register(ctx, connectReq)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Msg.Username, req.Username)
	assert.Equal(t, response.Msg.Name, req.Name)
	assert.Equal(t, response.Msg.Profile, req.Profile)
	assert.Equal(t, response.Msg.Premium, req.Premium)
	assert.Equal(t, response.Msg.Status, newUser.Status)
	assert.Equal(t, response.Msg.Gender, userv1.GENDER(userv1.GENDER_value[newUser.Gender]))
	mockRepoUser.AssertExpectations(t)
}
func testServiceAuthRegister_UserAlreadyExist(t *testing.T) {
	var mockRepoUser = new(mocks.RepositoryUser)
	var srv = NewService(&logrus.Logger{}, mockRepoUser)
	ctx := context.Background()

	req := authv1.RegisterRequest{
		Username: "testuser",
		Password: "testpassword",
		Gender:   userv1.GENDER_FEMALE,
		Name:     "testname",
		Profile:  "testprofile",
		Premium:  false,
	}
	connectReq := &connect.Request[authv1.RegisterRequest]{
		Msg: &req,
	}
	userResponse := model.User{
		ID:        uuid.NewString(),
		Username:  req.Username,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		DeletedAt: nil,
		Password:  "$2a$04$.j8bw6aqScH/CeVXetnDQ.5s/6NZZubXxKrNAbthotJ2e7jl.8BXO",
		Name:      req.Name,
		Profile:   req.Profile,
		Status:    "active",
		IsPremium: req.Premium,
		Gender:    userv1.GENDER_name[int32(req.Gender)],
	}

	mockRepoUser.On("DetailByUsername", ctx, connectReq.Msg.Username).Return(userResponse, nil).Once()
	response, err := srv.Register(ctx, connectReq)
	assert.Error(t, err)
	connectError := new(connect.Error)
	assert.ErrorAs(t, err, &connectError)
	assert.ErrorIs(t, err, constant.ErrUsernameExist)
	assert.Nil(t, response)
	mockRepoUser.AssertExpectations(t)
}
func testServiceAuthRegister_InternalServerError(t *testing.T) {
	var mockRepoUser = new(mocks.RepositoryUser)
	var srv = NewService(&logrus.Logger{}, mockRepoUser)
	ctx := context.Background()

	req := authv1.RegisterRequest{
		Username: "testuser",
		Password: "testpassword",
		Gender:   userv1.GENDER_FEMALE,
		Name:     "testname",
		Profile:  "testprofile",
		Premium:  false,
	}
	connectReq := &connect.Request[authv1.RegisterRequest]{
		Msg: &req,
	}
	mockRepoUser.On("DetailByUsername", ctx, connectReq.Msg.Username).Return(model.User{}, gorm.ErrInvalidDB).Once()
	response, err := srv.Register(ctx, connectReq)
	assert.Error(t, err)
	connectError := new(connect.Error)
	assert.ErrorAs(t, err, &connectError)
	assert.ErrorIs(t, err, constant.ErrInternalServer)
	assert.Nil(t, response)
	mockRepoUser.AssertExpectations(t)
}
