package create

import (
	"context"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
)

func Post(serverURL *url.URL, mongoURI string) {
	log.Println("post albums/create")

	ctx := context.Background()
	collection := MongoClient.Database("gostream").Collection("albums")

	filter := bson.M{"_id": CreatedID}
	result, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		log.Fatalf("failed to delete document: %s", err)
	}

	if result.DeletedCount != 1 {
		log.Fatalf("unexpected result during test cleanup")
	}

	collection = MongoClient.Database("gostream").Collection("tracks")
	filter = bson.M{"_id": "6f2b2305-2f65-461f-9008-da2cf46eac5c"}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatalf("failed to delete document: %s", err)
	}

	filter = bson.M{"_id": "b806c773-d9be-4e98-b241-a1a939e746d5"}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatalf("failed to delete document: %s", err)
	}

	err = MongoClient.Disconnect(ctx)
	if err != nil {
		log.Fatalf("failed to disconnect from database: %s", err)
	}
}
