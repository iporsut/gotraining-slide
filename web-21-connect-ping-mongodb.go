package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// Mongo URI pattern => mongodb://user:pass@host:port
	mongoURI := "mongodb://user:pass@localhost:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
