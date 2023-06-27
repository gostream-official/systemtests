package get

import (
	"context"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func Setup(serverURL *url.URL, mongoURI string) {
	log.Println("setup tracks/get")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	var err error
	ctx := context.Background()

	MongoClient, err = mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	artistToCreate := TrackInfo{
		ID:    "e28ad81a-05e9-41b3-8e37-e15176f4ae4e",
		Title: "A Random Track",
	}

	collection := MongoClient.Database("gostream").Collection("tracks")

	_, err = collection.InsertOne(ctx, artistToCreate)
	if err != nil {
		log.Fatalf("setup failed: %s", err)
	}
}
