package repository

import (
	"context"

	"github.com/felipecveiga/crud-puro-go/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	DB *mongo.Client
}

func NewUserRepository(clientDB *mongo.Client) *UserRepository {
	return &UserRepository{
		DB: clientDB,
	}
}

func (r *UserRepository) CreateUserDB(payload *model.User) error {
	coll := r.DB.Database("estudo_mongo").Collection("funcionarios")
	doc := model.User{
		ID:    payload.ID,
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
