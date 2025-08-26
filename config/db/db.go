package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection() *mongo.Client {

	db, err := mongo.Connect(context.Background(), (options.Client().ApplyURI("mongodb://localhost:27017")))
	if err != nil {
		log.Panicf("error connecting to database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
