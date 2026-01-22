package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/zmb3/spotify/v2"
)

type PlaylistInfo struct {
	ID   spotify.ID `json:"id"`
	Name string     `json:"name"`
}

func FetchUserPlaylists(ctx context.Context, client *spotify.Client) ([]byte, error) {
	// Fetch the first page; subsequent pages are pulled via NextPage.
	playlists, err := client.CurrentUsersPlaylists(ctx, spotify.Limit(50))
	// Collect only the fields we care about for downstream use.
	infos := []PlaylistInfo{}
	if err != nil {
		return nil, err
	}

	for {
		// Each page contains a slice of SimplePlaylist items.
		for _, p := range playlists.Playlists {
			infos = append(infos, PlaylistInfo{ID: p.ID, Name: p.Name})
		}
		// Advance to the next page until Spotify reports no more pages.
		err = client.NextPage(ctx, playlists)
		if err == spotify.ErrNoMorePages {
			break
		}
		if err != nil {
			return nil, err
		}
	}

	return json.Marshal(infos)

}

type SavedTrackInfo struct {
	ID         spotify.ID             `json:"id"`
	Name       string                 `json:"name"`
	Artists    []spotify.SimpleArtist `json:"artists"`
	DurationMs int                    `json:"duration_ms"`
	Popularity int                    `json:"popularity"`
}

func FetchLikedSongs(ctx context.Context, client *spotify.Client) error {
	page, err := client.CurrentUsersTracks(ctx, spotify.Limit(50))
	if err != nil {
		return err
	}

	file, err := os.Create("saved_tracks.ndjson")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	for {
		for _, p := range page.Tracks {
			entry := SavedTrackInfo{
				ID:         p.ID,
				Name:       p.Name,
				Artists:    p.Artists,
				DurationMs: int(p.Duration),
				Popularity: int(p.Popularity),
			}

			if err := encoder.Encode(entry); err != nil {
				return err
			}
		}

		err = client.NextPage(ctx, page)
		if err == spotify.ErrNoMorePages {
			break
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func FetchPlaylistItems() {
	//tracks, err := client.GetPlaylistItems(
	//	ctx,
	//	spotify.ID(playlistID),

}
