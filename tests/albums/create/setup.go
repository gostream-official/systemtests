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
	log.Println("setup albums/create")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	var err error
	ctx := context.Background()

	MongoClient, err = mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	trackInfoOne := TrackInfo{
		ID: "6f2b2305-2f65-461f-9008-da2cf46eac5c",
	}

	trackInfoTwo := TrackInfo{
		ID: "b806c773-d9be-4e98-b241-a1a939e746d5",
	}

	collection := MongoClient.Database("gostream").Collection("tracks")

	_, err = collection.InsertOne(ctx, trackInfoOne)
	if err != nil {
		log.Fatalf("setup failed: %s", err)
	}

	_, err = collection.InsertOne(ctx, trackInfoTwo)
	if err != nil {
		log.Fatalf("setup failed: %s", err)
	}
}
