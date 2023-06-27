package get

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/gostream-official/systemtests/pkg/assert"
)

func Test(serverURL *url.URL, mongoURI string) {
	log.Println("test artists/get")

	serverURL.Path = path.Join(serverURL.Path, "/artists", "/bdb0dc83-6a04-4e5d-ad70-7c7e9e0b21c4")
	request, err := http.NewRequest("GET", serverURL.String(), nil)

	if err != nil {
		log.Fatalf("failed to create http request")
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("failed to perform server request: %s", err)
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %s", err)
	}

	var unmarshaledResponse ArtistInfo
	err = json.Unmarshal(responseBody, &unmarshaledResponse)

	if err != nil {
		log.Fatalf("failed to unmarshal response body: %s", err)
	}

	assert.Expect(response.StatusCode == http.StatusOK, "expected the response status code to be 200")

	assert.Expect(unmarshaledResponse.ID == "bdb0dc83-6a04-4e5d-ad70-7c7e9e0b21c4", "expected the response id to be the same")
	assert.Expect(unmarshaledResponse.Name == "Martin Garrix", "expected the response name to be the same")
	assert.Expect(unmarshaledResponse.Genres[0] == "EDM", "expected the response genre to be the same")
	assert.Expect(unmarshaledResponse.Genres[1] == "Progressive House", "expected the response genre be the same")
	assert.Expect(unmarshaledResponse.Genres[2] == "Future House", "expected the response genre to be the same")
	assert.Expect(unmarshaledResponse.Followers == 3, "expected the response followers to be the same")
	assert.Expect(unmarshaledResponse.Stats.Popularity == 0.75, "expected the response popularity to be the same")
}
