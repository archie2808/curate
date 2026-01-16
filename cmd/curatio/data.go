package main

import (
	"context"

	"github.com/zmb3/spotify/v2"
)

type PlaylistInfo struct {
	ID   spotify.ID
	Name string
}

func FetchUserPlaylists(ctx context.Context, client *spotify.Client) ([]PlaylistInfo, error) {
	// Fetch the first page; subsequent pages are pulled via NextPage.
	playlists, err := client.CurrentUsersPlaylists(ctx)
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

	return infos, nil

}

func FetchPlaylistItems() {
	//tracks, err := client.GetPlaylistItems(
	//	ctx,
	//	spotify.ID(playlistID),

}
