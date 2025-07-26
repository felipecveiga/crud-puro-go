package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Sexo      string             `bson:"sexo"`
	Age       int                `bson:"age"`
	Phone     int                `bson:"phone"`
	Residence Residence          `bson:"residence,omitempty"`
}

type Residence struct {
	Street  string `bson:"street,omitempty"`
	City    string `bson:"city,omitempty"`
	Country string `bson:"country,omitempty"`
	Number  int    `bson:"number,omitempty"`
}
