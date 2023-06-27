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
	log.Println("setup albums/delete")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	var err error
	ctx := context.Background()

	MongoClient, err = mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	albumToCreate := AlbumInfo{
		ID:    "05577461-fdbe-4617-8f08-76c19d73b68b",
		Title: "A Random Album",
	}

	collection := MongoClient.Database("gostream").Collection("albums")

	_, err = collection.InsertOne(ctx, albumToCreate)
	if err != nil {
		log.Fatalf("setup failed: %s", err)
	}
}
