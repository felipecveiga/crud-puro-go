package service

import (
	"errors"
	"testing"

	"github.com/felipecveiga/crud-puro-go/errs"
	"github.com/felipecveiga/crud-puro-go/model"
	"github.com/felipecveiga/crud-puro-go/repository"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestCreateUser_WhenCreateAccount_ReturnSucess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepository(ctrl)
	service := NewUserService(mockRepo)

	payload := &model.User{
		Name:  "Felipe",
		Email: "felipe@gmail.com",
		Sexo:  "M",
		Age:   31,
		Phone: 21212121,
		Residence: model.Residence{
			Street:  "rua A",
			City:    "rio de janeiro",
			Country: "Brasil",
			Number:  27,
		},
	}

	mockRepo.EXPECT().
		CreateUserDB(payload).
		Return(nil)

	err := service.CreateUser(payload)

	assert.NoError(t, err)
}

func TestCreateUser_WhenCreateAccount_ReturnErrorNameEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepository(ctrl)
	service := NewUserService(mockRepo)

	payload := &model.User{
		Name:  "",
		Email: "felipe@gmail.com",
		Sexo:  "M",
		Age:   31,
		Phone: 21212121,
		Residence: model.Residence{
			Street:  "rua A",
			City:    "rio de janeiro",
			Country: "Brasil",
			Number:  27,
		},
	}

	err := service.CreateUser(payload)

	assert.Error(t, err)
	assert.EqualError(t, err, "missing required fields: name, email, and phone number")
}

func TestCreateUser_WhenCreateAccount_ReturnErrorCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepository(ctrl)
	service := NewUserService(mockRepo)

	payload := &model.User{
		Name:  "Felipe",
		Email: "felipe@gmail.com",
		Sexo:  "M",
		Age:   31,
		Phone: 21212121,
		Residence: model.Residence{
			Street:  "rua A",
			City:    "rio de janeiro",
			Country: "Brasil",
			Number:  27,
		},
	}

	mockRepo.EXPECT().
		CreateUserDB(payload).
		Return(errors.New("some error"))

	err := service.CreateUser(payload)

	assert.Error(t, err)
	assert.EqualError(t, err, "some error")
}

func TestGetUser_WhenGetUser_ReturnSucess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepository(ctrl)
	service := NewUserService(mockRepo)

	user := &model.User{
		Name:  "Felipe",
		Email: "felipe@gmail.com",
		Sexo:  "M",
		Age:   31,
		Phone: 21212121,
		Residence: model.Residence{
			Street:  "rua A",
			City:    "rio de janeiro",
			Country: "Brasil",
			Number:  27,
		},
	}

	id := "68a8e66a5a3b238655f42f4"
	mockRepo.EXPECT().
		FindByID(id).
		Return(user, nil)

	result, err := service.GetUser(id)

	assert.NoError(t, err)
	assert.Equal(t, user, result)
}

func TestGetUser_WhenGetUser_ReturnErrorUserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepository(ctrl)
	service := NewUserService(mockRepo)

	id := "68a8e66a5a3b238655f42f4"
	mockRepo.EXPECT().
		FindByID(id).
		Return(nil, errs.ErrUserNotFound)

	_, err := service.GetUser(id)

	assert.Error(t, err)
	assert.EqualError(t, err, "user not found")
}

func TestGetUser_WhenGetUser_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepository(ctrl)
	service := NewUserService(mockRepo)

	id := "68a8e66a5a3b238655f42f4"
	mockRepo.EXPECT().
		FindByID(id).
		Return(nil, errs.ErrUserSearchFailed)

	_, err := service.GetUser(id)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to search for user")
}