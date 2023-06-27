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
	log.Println("test albums/get")

	serverURL.Path = path.Join(serverURL.Path, "/albums", "/231f02d7-161d-43ae-b2e6-9c58b636959c")
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

	var unmarshaledResponse AlbumInfo
	err = json.Unmarshal(responseBody, &unmarshaledResponse)

	if err != nil {
		log.Fatalf("failed to unmarshal response body: %s", err)
	}

	assert.Expect(response.StatusCode == http.StatusOK, "expected the response status code to be 200")

	assert.Expect(unmarshaledResponse.ID == "231f02d7-161d-43ae-b2e6-9c58b636959c", "expected the response id to be the same")
	assert.Expect(unmarshaledResponse.Title == "A Random Album", "expected the response title to be the same")
	assert.Expect(unmarshaledResponse.TrackIDs[0] == "a", "expected the response trackID to be the same")
	assert.Expect(unmarshaledResponse.TrackIDs[1] == "b", "expected the response trackID be the same")
	assert.Expect(unmarshaledResponse.TrackIDs[2] == "c", "expected the response trackID to be the same")
	assert.Expect(unmarshaledResponse.Stats.Popularity == 0.75, "expected the response popularity to be the same")
}
