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
	log.Print("test tracks/update")

	updateArtistInfo := UpdateTrackInfo{
		Title: "Dave's Track",
	}

	marshaledBody, err := json.Marshal(updateArtistInfo)
	if err != nil {
		log.Fatalf("failed to marshal request body: %s", err)
	}

	serverURL.Path = path.Join(serverURL.Path, "/tracks", "/86bb41b0-ab2a-4a95-8d05-ba05e996cacc")
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
	collection := MongoClient.Database("gostream").Collection("tracks")

	var result TrackInfo

	filter := bson.M{"_id": "86bb41b0-ab2a-4a95-8d05-ba05e996cacc"}
	err = collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		log.Fatalf("failed to find or decode document: %s", err)
	}

	assert.Expect(result.Title == updateArtistInfo.Title, "expected track title to be the same")
}
