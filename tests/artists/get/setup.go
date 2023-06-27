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
	log.Println("setup artists/get")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	var err error
	ctx := context.Background()

	MongoClient, err = mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	artistToCreate := ArtistInfo{
		ID:        "bdb0dc83-6a04-4e5d-ad70-7c7e9e0b21c4",
		Name:      "Martin Garrix",
		Genres:    []string{"EDM", "Progressive House", "Future House"},
		Followers: 3,
		Stats: ArtistStats{
			Popularity: 0.75,
		},
	}

	collection := MongoClient.Database("gostream").Collection("artists")

	_, err = collection.InsertOne(ctx, artistToCreate)
	if err != nil {
		log.Fatalf("setup failed: %s", err)
	}
}
