package update

import (
	"context"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func Setup(serverURL *url.URL, mongoURI string) {
	log.Println("setup tracks/update")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	var err error
	ctx := context.Background()

	MongoClient, err = mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	artistToCreate := TrackInfo{
		ID:    "86bb41b0-ab2a-4a95-8d05-ba05e996cacc",
		Title: "A Random Track",
	}

	collection := MongoClient.Database("gostream").Collection("tracks")

	_, err = collection.InsertOne(ctx, artistToCreate)
	if err != nil {
		log.Fatalf("setup failed: %s", err)
	}
}
