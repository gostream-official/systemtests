package delete

import (
	"context"
	"log"
	"net/url"
)

func Post(serverURL *url.URL, mongoURI string) {
	log.Println("post artists/delete")

	ctx := context.Background()
	err := MongoClient.Disconnect(ctx)

	if err != nil {
		log.Fatalf("failed to disconnect from database: %s", err)
	}
}
