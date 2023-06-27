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
	log.Println("test tracks/create")

	trackToCreate := CreateTrackInfo{
		ArtistID:          "736ed31c-180d-49f7-aac5-5867a01727ee",
		FeaturedArtistIDs: []string{"403a6108-0e76-4534-b12b-5beecf69e3d1"},
		Title:             "Demo Track",
		Label:             "Demo Label",
		ReleaseDate:       "2023-06-02",
		TrackStats: TrackStats{
			Streams: 100000,
			Likes:   3000,
		},
		AudioFeatures: AudioFeatures{
			Key: "F# Major",
		},
	}

	marshaledBody, err := json.Marshal(trackToCreate)
	if err != nil {
		log.Fatalf("failed to marshal request body: %s", err)
	}

	serverURL.Path = path.Join(serverURL.Path, "/tracks")
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

	var unmarshaledResponse TrackInfo
	err = json.Unmarshal(responseBody, &unmarshaledResponse)

	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %s", err)
	}

	CreatedID = unmarshaledResponse.ID

	assert.Expect(response.StatusCode == 200, "expected the response status code to be 200")
	assert.Expect(unmarshaledResponse.ArtistID == trackToCreate.ArtistID, "expected response artist id to be the same")
	assert.Expect(unmarshaledResponse.FeaturedArtistIDs[0] == trackToCreate.FeaturedArtistIDs[0], "expected response first featured artist to be the same")
	assert.Expect(unmarshaledResponse.Title == trackToCreate.Title, "expected response title to be the same")
	assert.Expect(unmarshaledResponse.Label == trackToCreate.Label, "expected response label to be the same")
	assert.Expect(unmarshaledResponse.ReleaseDate.Format("2006-01-02") == trackToCreate.ReleaseDate, "expected response release date to be the same")

	assert.Expect(unmarshaledResponse.TrackStats.Streams == trackToCreate.TrackStats.Streams, "expected response streams to be the same")
	assert.Expect(unmarshaledResponse.TrackStats.Likes == trackToCreate.TrackStats.Likes, "expected response likes to be the same")

	assert.Expect(unmarshaledResponse.AudioFeatures.Key == trackToCreate.AudioFeatures.Key, "expected response key to be the same")
	assert.Expect(unmarshaledResponse.AudioFeatures.Tempo == trackToCreate.AudioFeatures.Tempo, "expected response tempo to be the same")
	assert.Expect(unmarshaledResponse.AudioFeatures.Duration == trackToCreate.AudioFeatures.Duration, "expected response duration to be the same")
	assert.Expect(unmarshaledResponse.AudioFeatures.Energy == trackToCreate.AudioFeatures.Energy, "expected response energy to be the same")
	assert.Expect(unmarshaledResponse.AudioFeatures.Danceability == trackToCreate.AudioFeatures.Danceability, "expected response danceability to be the same")
	assert.Expect(unmarshaledResponse.AudioFeatures.Accousticness == trackToCreate.AudioFeatures.Accousticness, "expected response accousticness to be the same")
	assert.Expect(unmarshaledResponse.AudioFeatures.Instrumentalness == trackToCreate.AudioFeatures.Instrumentalness, "expected response instrumentalness to be the same")
	assert.Expect(unmarshaledResponse.AudioFeatures.Liveness == trackToCreate.AudioFeatures.Liveness, "expected response liveness to be the same")
	assert.Expect(unmarshaledResponse.AudioFeatures.Loudness == trackToCreate.AudioFeatures.Loudness, "expected response loudness to be the same")
	assert.Expect(unmarshaledResponse.AudioFeatures.TimeSignature == trackToCreate.AudioFeatures.TimeSignature, "expected response time signature to be the same")
}
