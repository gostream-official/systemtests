package main

import (
	"flag"
	"log"
	"net/url"

	"github.com/gostream-official/systemtests/tests/tracks/delete"
)

func main() {
	var flagServerURL string
	var flagMongoConnectionURI string

	flag.StringVar(&flagServerURL, "server-url", "http://127.0.0.1:9871/", "Specifies the server url")
	flag.StringVar(&flagMongoConnectionURI, "mongo-uri", "", "Specifies the MongoDB connection uri")

	flag.Parse()

	serverURL, err := url.Parse(flagServerURL)
	if err != nil {
		log.Fatalf("failed to parse server url")
	}

	mongoURI := flagMongoConnectionURI

	delete.Setup(serverURL, mongoURI)
	delete.Test(serverURL, mongoURI)
	delete.Post(serverURL, mongoURI)
}
