package service

import (
	"errors"

	"github.com/felipecveiga/crud-puro-go/model"
	"github.com/felipecveiga/crud-puro-go/repository"
)

type UserService struct {
	Repository *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{
		Repository: r,
	}
}

func (s *UserService) CreateUser(payload *model.User) error {

	if payload.Name == "" || payload.Email == "" || payload.Phone == 0 {
		return errors.New("erro ao criar conta, preenchimento obrigatório do nome, email e telefone")
	}

	err := s.Repository.CreateUserDB(payload)
	if err != nil {
		return err
	}

	return nil
}
