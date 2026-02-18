package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

//Spotify oauth



// Redirect URI for spotify API
const redirectURI = "https://127.0.0.1:8888/callback"


// TLS config
const keyFile = "/Users/ag1/Documents/curatio_httpscert/127.0.0.1-key.pem"
const certFile = "/Users/ag1/Documents/curatio_httpscert/127.0.0.1.pem"

var (
	// OAuth client config for Spotify.
	auth = spotifyauth.New(
		spotifyauth.WithClientID(os.Getenv("SPOTIFY_CLIENT_ID")),
		spotifyauth.WithClientSecret(os.Getenv("SPOTIFY_CLIENT_SECRET")),
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserReadPrivate,
			spotifyauth.ScopePlaylistReadPrivate,
			spotifyauth.ScopePlaylistReadCollaborative,
			spotifyauth.ScopeUserLibraryRead,
		),
	)

	// Channel used to hand the authenticated client from the callback to main.
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

func Login(certFile, keyFile string) (*spotify.Client, error) {
      // Register the callback handler (uses completeAuth internally)
      http.HandleFunc("/callback", completeAuth)
      http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
          log.Println("Got request for:", r.URL.String())
      })

      // Start HTTPS server
      go http.ListenAndServeTLS(":8888", certFile, keyFile, nil)

      // Print login URL
      url := auth.AuthURL(state)
      fmt.Println("Please log into spotify by visiting the following page:", url)

      // Block until callback sends the client
      client := <- ch
      return client, nil
  }

func completeAuth(w http.ResponseWriter, r *http.Request) {


	// Exchange the callback code for an access token.
	tok, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "couldnt get token", http.StatusForbidden)
		log.Fatal(err)
	}

	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// Use the token to get an authenticated client.
	client := spotify.New(auth.Client(r.Context(), tok))
	fmt.Fprintf(w, "Login Completed!")
	ch <- client
}
