package delete

import (
	"context"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func Setup(serverURL *url.URL, mongoURI string) {
	log.Println("setup tracks/delete")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	var err error
	ctx := context.Background()

	MongoClient, err = mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	albumToCreate := TrackInfo{
		ID:    "9557a853-f848-49dd-a48d-2fd0d38b6568",
		Title: "A Random Track",
	}

	collection := MongoClient.Database("gostream").Collection("tracks")

	_, err = collection.InsertOne(ctx, albumToCreate)
	if err != nil {
		log.Fatalf("setup failed: %s", err)
	}
}
