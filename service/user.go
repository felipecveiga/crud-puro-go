package service

import (
	"github.com/felipecveiga/crud-puro-go/errs"
	"github.com/felipecveiga/crud-puro-go/model"
	"github.com/felipecveiga/crud-puro-go/repository"
)

//go:generate mockgen -source=./user.go -destination=./user_mock.go -package=service
type Service interface {
	CreateUser(payload *model.User) error
	GetUser(id string) (*model.User, error)
}

type service struct {
	Repository repository.Repository
}

func NewUserService(r repository.Repository) Service {
	return &service{
		Repository: r,
	}
}

func (s *service) CreateUser(payload *model.User) error {

	if payload.Name == "" || payload.Email == "" || payload.Phone == 0 {
		return errs.ErrMissingRequiredFields
	}

	err := s.Repository.CreateUserDB(payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetUser(id string) (*model.User, error) {

	user, err := s.Repository.FindByID(id)
	if err != nil {
		return &model.User{}, err
	}

	return user, nil
}
