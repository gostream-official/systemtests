package create

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/gostream-official/systemtests/pkg/assert"
)

var CreatedID string

func Test(serverURL *url.URL, mongoURI string) {
	log.Println("test artists/create")

	artistToCreate := CreateArtistInfo{
		Name:      "Martin Garrix",
		Genres:    []string{"EDM", "Progressive House", "Future House"},
		Followers: 3,
		Stats: ArtistStats{
			Popularity: 0.75,
		},
	}

	marshaledBody, err := json.Marshal(artistToCreate)
	if err != nil {
		log.Fatalf("failed to marshal request body: %s", err)
	}

	serverURL.Path = path.Join(serverURL.Path, "/artists")
	request, err := http.NewRequest("POST", serverURL.String(), bytes.NewBuffer(marshaledBody))

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

	CreatedID = unmarshaledResponse.ID

	assert.Expect(response.StatusCode == http.StatusOK, "expected the response status code to be 200")

	assert.Expect(unmarshaledResponse.Name == artistToCreate.Name, "expected response name to be the same")
	assert.Expect(unmarshaledResponse.Genres[0] == artistToCreate.Genres[0], "expected response genre[0] to be the same")
	assert.Expect(unmarshaledResponse.Genres[1] == artistToCreate.Genres[1], "expected response genre[0] to be the same")
	assert.Expect(unmarshaledResponse.Genres[2] == artistToCreate.Genres[2], "expected response genre[0] to be the same")
	assert.Expect(unmarshaledResponse.Followers == artistToCreate.Followers, "expected response followers to be the same")

	assert.Expect(unmarshaledResponse.Stats.Popularity == artistToCreate.Stats.Popularity, "expected response popularity to be the same")
}
