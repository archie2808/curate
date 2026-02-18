package main

import (
	"context"
	"curatio/internal/auth"
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

	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Pass UID to playlist creation in curator flow when we get to that point
	userID := user.ID

	likedSongs, err := spotify.FetchLikedSongs(context.Background(), client)
	if err != nil {
		log.Fatal(err)
	}

}
