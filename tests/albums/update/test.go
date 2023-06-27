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
	log.Print("test albums/update")

	updateArtistInfo := UpdateAlbumInfo{
		Title: "Dave's Album",
	}

	marshaledBody, err := json.Marshal(updateArtistInfo)
	if err != nil {
		log.Fatalf("failed to marshal request body: %s", err)
	}

	serverURL.Path = path.Join(serverURL.Path, "/albums", "/0a9e0aba-8cda-444a-83a1-d0eed25955cb")
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
	collection := MongoClient.Database("gostream").Collection("albums")

	var result AlbumInfo

	filter := bson.M{"_id": "0a9e0aba-8cda-444a-83a1-d0eed25955cb"}
	err = collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		log.Fatalf("failed to find or decode document: %s", err)
	}

	assert.Expect(result.Title == updateArtistInfo.Title, "expected album title to be the same")
}
