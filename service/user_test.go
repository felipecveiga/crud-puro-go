package service

import (
	"errors"
	"testing"

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
	assert.EqualError(t, err, "erro ao criar conta, preenchimento obrigatório do nome, email e telefone")
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
