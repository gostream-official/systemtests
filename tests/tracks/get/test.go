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
	log.Println("test tracks/get")

	serverURL.Path = path.Join(serverURL.Path, "/tracks", "/e28ad81a-05e9-41b3-8e37-e15176f4ae4e")
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

	var unmarshaledResponse TrackInfo
	err = json.Unmarshal(responseBody, &unmarshaledResponse)

	if err != nil {
		log.Fatalf("failed to unmarshal response body: %s", err)
	}

	assert.Expect(response.StatusCode == http.StatusOK, "expected the response status code to be 200")

	assert.Expect(unmarshaledResponse.ID == "e28ad81a-05e9-41b3-8e37-e15176f4ae4e", "expected the response id to be the same")
	assert.Expect(unmarshaledResponse.Title == "A Random Track", "expected the response title to be the same")
}
