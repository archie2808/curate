package spotify

import (
	"context"
	"log"

	"github.com/zmb3/spotify/v2"
)

// FetchLikedSongs && FetchArtistGenres

type ArtistInfo struct {
	ID     spotify.ID `json:"id"`
	Name   string     `json:"name"`
	Genres []string   `json:"genres"`
}

type SavedTrackInfo struct {
	ID         spotify.ID   `json:"id"`
	Name       string       `json:"name"`
	Artists    []ArtistInfo `json:"artists"`
	DurationMs int          `json:"duration_ms"`
}

func FetchTrackInfo(ctx context.Context, client *spotify.Client) ([]SavedTrackInfo, error) {
	page, err := client.CurrentUsersTracks(ctx, spotify.Limit(50))
	if err != nil {
		return nil, err
	}

	trackInfo := make([]SavedTrackInfo, 0)

	for {
		for _, p := range page.Tracks {
			artists := make([]ArtistInfo, len(p.Artists))
			for i, a := range p.Artists {
				artists[i] = ArtistInfo{
					ID:   a.ID,
					Name: a.Name,
					// Genres filled in by enrichArtistsGenre
				}

				trackInfo = append(trackInfo, SavedTrackInfo{
					ID:         p.ID,
					Name:       p.Name,
					Artists:    artists,
					DurationMs: int(p.Duration),
				})

			}

		}

		err = client.NextPage(ctx, page)
		if err == spotify.ErrNoMorePages {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	err = enrichArtistGenres(ctx, client, trackInfo)
	if err != nil {
		return nil, err
	}
	return trackInfo, nil
	
}

func enrichArtistGenres(ctx context.Context, client *spotify.Client, trackInfo []SavedTrackInfo) error {
	// Collect unique artist IDs
	// Batch getArtist calls
	// Need to access full artist to get genre by of a given artist by ID

	// Populate only keys with artistID (removes dupes)
	artistIDSet := make(map[spotify.ID]struct{})
	for _, trackStruct := range trackInfo {

		for _, artistStruct := range trackStruct.Artists {
			artistIDSet[artistStruct.ID] = struct{}{}
		}
	}
	// Convert back to slice for giving to GetArtists
	artistIDs := make([]spotify.ID, 0, len(artistIDSet))
	for key := range artistIDSet {
		artistIDs = append(artistIDs, key)
	}
	// GetArtist expects single spotify.ID not []spotify.ID so need to range over it

	fullArtist, err := client.GetArtists(ctx, artistIDs...)
	if err != nil {
		return err
	}
	log.Print(fullArtist)
	return nil 
}
