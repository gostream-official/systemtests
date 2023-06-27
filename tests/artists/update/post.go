package update

import (
	"context"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
)

func Post(serverURL *url.URL, mongoURI string) {
	log.Println("post artists/update")

	ctx := context.Background()
	collection := MongoClient.Database("gostream").Collection("artists")

	filter := bson.M{"_id": "d5f27ff4-71e1-47d1-bf29-29e51e7145bb"}
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
