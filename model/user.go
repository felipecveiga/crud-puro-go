package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Sexo      string             `bson:"sexo,omitempty"`
	Age       int                `bson:"age,omitempty"`
	Phone     int                `bson:"phone,omitempty"`
	Residence Residence          `bson:"residence,omitempty"`
}

type Residence struct {
	Street  string `bson:"street,omitempty"`
	City    string `bson:"city,omitempty"`
	Country string `bson:"country,omitempty"`
	Number  int    `bson:"number,omitempty"`
}
