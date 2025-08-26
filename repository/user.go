package repository

import (
	"context"
	"errors"

	"github.com/felipecveiga/crud-puro-go/model"

	"github.com/felipecveiga/crud-puro-go/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -source=./user.go -destination=./user_mock.go -package=repository

type Repository interface {
	CreateUserDB(payload *model.User) error
	FindByID(id string) (*model.User, error)
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
		return errs.ErrInsertUserDatabase
	}

	return nil
}

func (r *repository) FindByID(id string) (*model.User, error) {
	coll := r.DB.Database("estudo_mongo").Collection("funcionarios")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errs.ErrConvertIDObjectID
	}

	filter := bson.D{{Key: "_id", Value: objectID}}

	var user model.User
	err = coll.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrUserNotFound
		}

		return nil, errs.ErrSearchUser
	}

	return &user, nil
}
