package model

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID        bson.ObjectID `bson:"_id"`
	Name      string        `bson:"name"`
	Email     string        `bson:"email"`
	Sexo      string        `bson:"sexo"`
	Age       int           `bson:"age"`
	Phone     int           `bson:"phone"`
	Residence Residence     `bson:"residence"`
}

type Residence struct {
	Street  string `bson:"street"`
	City    string `bson:"city"`
	Country string `bson:"country"`
	Number  int    `bson:"number"`
}
