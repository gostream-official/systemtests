package main

import (
	"flag"
	"log"
	"net/url"

	"github.com/gostream-official/systemtests/tests/albums/create"
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

	create.Setup(serverURL, mongoURI)
	create.Test(serverURL, mongoURI)
	create.Post(serverURL, mongoURI)
}
