package repository

import (
	"context"

	"github.com/felipecveiga/crud-puro-go/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

//go:generate mockgen -source=./user.go -destination=./user_mock.go -package=repository

type Repository interface {
	CreateUserDB(payload *model.User) error
}

type repository struct {
	DB *mongo.Client
}

func NewUserRepository(clientDB *mongo.Client) Repository {
	return &repository{
		DB: clientDB,
	}
}

func (r *repository) CreateUserDB(payload *model.User) error {
	coll := r.DB.Database("estudo_mongo").Collection("funcionarios")
	doc := model.User{
		Name:  payload.Name,
		Email: payload.Email,
		Sexo:  payload.Sexo,
		Age:   payload.Age,
		Phone: payload.Phone,
		Residence: model.Residence{
			Street:  payload.Residence.Street,
			City:    payload.Residence.City,
			Country: payload.Residence.Country,
			Number:  payload.Residence.Number,
		},
	}

	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}

	return nil
}
