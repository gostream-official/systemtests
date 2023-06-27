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
	log.Println("setup artists/delete")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	var err error
	ctx := context.Background()

	MongoClient, err = mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	artistToCreate := ArtistInfo{
		ID:        "6e371498-dc6a-4b3e-91e2-895285835b57",
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
