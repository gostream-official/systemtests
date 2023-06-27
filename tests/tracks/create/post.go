package create

import (
	"context"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
)

func Post(serverURL *url.URL, mongoURI string) {
	log.Println("post tracks/create")

	ctx := context.Background()
	collection := MongoClient.Database("gostream").Collection("tracks")

	filter := bson.M{"_id": CreatedID}
	result, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		log.Fatalf("failed to delete document: %s", err)
	}

	if result.DeletedCount != 1 {
		log.Fatalf("unexpected result during test cleanup")
	}

	collection = MongoClient.Database("gostream").Collection("artists")
	filter = bson.M{"_id": "736ed31c-180d-49f7-aac5-5867a01727ee"}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatalf("failed to delete document: %s", err)
	}

	filter = bson.M{"_id": "403a6108-0e76-4534-b12b-5beecf69e3d1"}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatalf("failed to delete document: %s", err)
	}

	err = MongoClient.Disconnect(ctx)
	if err != nil {
		log.Fatalf("failed to disconnect from database: %s", err)
	}
}
