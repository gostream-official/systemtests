package get

import (
	"context"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
)

func Post(serverURL *url.URL, mongoURI string) {
	log.Println("post albums/get")

	ctx := context.Background()
	collection := MongoClient.Database("gostream").Collection("albums")

	filter := bson.M{"_id": "231f02d7-161d-43ae-b2e6-9c58b636959c"}
	result, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		log.Fatalf("failed to delete document: %s", err)
	}

	if result.DeletedCount != 1 {
		log.Fatalf("unexpected result during test cleanup")
	}

	err = MongoClient.Disconnect(ctx)
	if err != nil {
		log.Fatalf("failed to disconnect from database: %s", err)
	}
}
