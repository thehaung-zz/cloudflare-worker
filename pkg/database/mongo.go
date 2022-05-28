package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func InitializeConnection(dbName string, mongoURI string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	log.Println(mongoURI)
	if err != nil {
		log.Fatalf("Error while connecting to mongoDB: %v\n", err)
	}
	return client.Database(dbName)
}
