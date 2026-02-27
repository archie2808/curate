package main

import (
	"context"
	"curatio/internal/auth"
	"curatio/internal/llm"
	"curatio/internal/profile"
	"curatio/internal/spotify"
	"log"
)

// TLS config
const keyFile = "/Users/ag1/Documents/curatio_httpscert/127.0.0.1-key.pem"
const certFile = "/Users/ag1/Documents/curatio_httpscert/127.0.0.1.pem"

// Redirect URI for spotify API
const redirectURI = "https://127.0.0.1:8888/callback"

func main() {
	client, err := auth.Login(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	trackInfo, err := spotify.FetchTrackInfo(context.Background(), client)
	if err != nil {
		log.Fatal(err)
	}

	artistSample := profile.GetArtistSample(trackInfo)
	config := llm.PromptInstructions()

	buckets := llm.GenreBuckets(artistSample, config)
	log.Println(buckets)
}
