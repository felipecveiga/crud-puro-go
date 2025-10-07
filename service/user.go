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
	GetAllUsers() ([]model.User, error)
	DeleteUser(id string) error
	UpdateUser(id string, payload *model.User) error
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

func (s *service) GetAllUsers() ([]model.User, error) {

	users, err := s.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *service) DeleteUser(id string) error {

	if len(id) != 24 {
		return errs.ErrIDInvalid
	}

	_, err := s.Repository.FindByID(id)
	if err != nil {
		return errs.ErrUserNotFound
	}

	_, err = s.Repository.DeleteUserByID(id)
	if err != nil {
		return errs.ErrDeleteUser
	}

	return nil
}

func (s *service) UpdateUser(id string, payload *model.User) error {
	if len(id) != 24 {
		return errs.ErrIDInvalid
	}

	_, err := s.Repository.FindByID(id)
	if err != nil {
		return errs.ErrUserNotFound
	}

	_, err = s.Repository.UpdateUserByID(id, payload)
	if err != nil {
		return errs.ErrDeleteUser
	}

	return nil
}
