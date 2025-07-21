package service

import (
	"errors"

	"github.com/felipecveiga/crud-puro-go/model"
)

func CreateUser(payload *model.User) (string, error) {

	if payload.Name == "" {
		return "", errors.New("erro ao criar conta")
	}

	return "create Sucess", nil
}
