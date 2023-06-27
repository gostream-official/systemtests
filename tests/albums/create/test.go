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

	albumToCreate := CreateAlbumInfo{
		Title: "A Random Album",
		TrackIDs: []string{
			"6f2b2305-2f65-461f-9008-da2cf46eac5c",
			"b806c773-d9be-4e98-b241-a1a939e746d5",
		},
		Stats: AlbumStats{
			Popularity: 0.75,
		},
	}

	marshaledBody, err := json.Marshal(albumToCreate)
	if err != nil {
		log.Fatalf("failed to marshal request body: %s", err)
	}

	serverURL.Path = path.Join(serverURL.Path, "/albums")
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

	var unmarshaledResponse AlbumInfo
	err = json.Unmarshal(responseBody, &unmarshaledResponse)

	if err != nil {
		log.Fatalf("failed to unmarshal response body: %s", err)
	}

	CreatedID = unmarshaledResponse.ID

	assert.Expect(response.StatusCode == http.StatusOK, "expected the response status code to be 200")

	assert.Expect(unmarshaledResponse.Title == albumToCreate.Title, "expected response title to be the same")
	assert.Expect(unmarshaledResponse.TrackIDs[0] == albumToCreate.TrackIDs[0], "expected response trackIDs[0] to be the same")
	assert.Expect(unmarshaledResponse.TrackIDs[1] == albumToCreate.TrackIDs[1], "expected response trackIDs[0] to be the same")
	assert.Expect(unmarshaledResponse.Stats.Popularity == albumToCreate.Stats.Popularity, "expected response popularity to be the same")
}
