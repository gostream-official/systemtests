package create

import (
	"context"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func Setup(serverURL *url.URL, mongoURI string) {
	log.Println("setup tracks/create")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	var err error
	ctx := context.Background()

	MongoClient, err = mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	artistInfoOne := ArtistInfo{
		ID: "736ed31c-180d-49f7-aac5-5867a01727ee",
	}

	artistInfoTwo := TrackInfo{
		ID: "403a6108-0e76-4534-b12b-5beecf69e3d1",
	}

	collection := MongoClient.Database("gostream").Collection("artists")

	_, err = collection.InsertOne(ctx, artistInfoOne)
	if err != nil {
		log.Fatalf("setup failed: %s", err)
	}

	_, err = collection.InsertOne(ctx, artistInfoTwo)
	if err != nil {
		log.Fatalf("setup failed: %s", err)
	}
}
