package update

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/gostream-official/systemtests/pkg/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test(serverURL *url.URL, mongoURI string) {
	log.Print("test artists/update")

	updateArtistInfo := UpdateArtistInfo{
		Name: "David Guetta",
	}

	marshaledBody, err := json.Marshal(updateArtistInfo)
	if err != nil {
		log.Fatalf("failed to marshal request body: %s", err)
	}

	serverURL.Path = path.Join(serverURL.Path, "/artists", "/d5f27ff4-71e1-47d1-bf29-29e51e7145bb")
	request, err := http.NewRequest("PUT", serverURL.String(), bytes.NewBuffer(marshaledBody))

	if err != nil {
		log.Fatalf("failed to create http request")
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("failed to perform server request: %s", err)
	}

	defer response.Body.Close()

	assert.Expect(response.StatusCode == http.StatusNoContent, "expected the response status code to be 204")

	ctx := context.Background()
	collection := MongoClient.Database("gostream").Collection("artists")

	var result ArtistInfo

	filter := bson.M{"_id": "d5f27ff4-71e1-47d1-bf29-29e51e7145bb"}
	err = collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		log.Fatalf("failed to find or decode document: %s", err)
	}

	assert.Expect(result.Name == updateArtistInfo.Name, "expected artist name to be the same")
}
