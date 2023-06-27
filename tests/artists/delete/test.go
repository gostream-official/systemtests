package delete

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/gostream-official/systemtests/pkg/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test(serverURL *url.URL, mongoURI string) {
	log.Println("test artists/delete")

	serverURL.Path = path.Join(serverURL.Path, "/artists", "/6e371498-dc6a-4b3e-91e2-895285835b57")
	request, err := http.NewRequest("DELETE", serverURL.String(), nil)

	if err != nil {
		log.Fatalf("failed to create http request")
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("failed to perform server request: %s", err)
	}

	defer response.Body.Close()

	assert.Expect(response.StatusCode == http.StatusAccepted, "expected the response status code to be 202")

	ctx := context.Background()
	collection := MongoClient.Database("gostream").Collection("artists")

	filter := bson.M{"_id": "6e371498-dc6a-4b3e-91e2-895285835b57"}
	count, err := collection.CountDocuments(ctx, filter)

	if err != nil {
		log.Fatalf("failed to count documents: %s", err)
	}

	assert.Expect(count == 0, "expected document count after deletion to be 0")
}
