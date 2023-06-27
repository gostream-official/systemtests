package update

import (
	"context"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
)

func Post(serverURL *url.URL, mongoURI string) {
	log.Println("post albums/update")

	ctx := context.Background()
	collection := MongoClient.Database("gostream").Collection("albums")

	filter := bson.M{"_id": "0a9e0aba-8cda-444a-83a1-d0eed25955cb"}
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
