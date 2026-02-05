package main

import (
	"context"

	"github.com/zmb3/spotify/v2"
)

type PlaylistInfo struct {
	ID   spotify.ID `json:"id"`
	Name string     `json:"name"`
}

/*func FetchUserPlaylists(ctx context.Context, client *spotify.Client) ([]byte, error) {
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

} */

type SavedTrackInfo struct {
	ID         spotify.ID             `json:"id"`
	Name       string                 `json:"name"`
	Artists    []spotify.SimpleArtist `json:"artists"`
	DurationMs int                    `json:"duration_ms"`
	Popularity int                    `json:"popularity"`
	/*Acousticness 	float32				   `json:"Acousticness"`
	Danceability	float32				   `json:"Danceability"`
	Energy			float32 			   `json:"Energy"`
	Valence			float32				   `json:"Valence"`
	Tempo			float32				   `json:"Tempo"` */
}

func FetchLikedSongs(ctx context.Context, client *spotify.Client) ([]SavedTrackInfo, error) {
	page, err := client.CurrentUsersTracks(ctx, spotify.Limit(50))
	if err != nil {
		return nil, err
	}

	infos := make([]SavedTrackInfo, 0)

	for {
		for _, p := range page.Tracks {
			infos = append(infos, SavedTrackInfo{
				ID:         p.ID,
				Name:       p.Name,
				Artists:    p.Artists,
				DurationMs: int(p.Duration),
				Popularity: int(p.Popularity),
			})

		}

		err = client.NextPage(ctx, page)
		if err == spotify.ErrNoMorePages {
			break
		}
		if err != nil {
			return nil, err
		}
	}

	return infos, nil
}

// We're still going to have to iterate over liked songs but for each item ID
// we shoud be extracting audio features and appending to struct
func FetchAudioCharacteristics(ctx context.Context, client *spotify.Client) (map[spotify.ID]*spotify.AudioFeatures, error) {
	savedTracks, err := FetchLikedSongs(ctx, client)
	if err != nil {
		return nil, err
	}

	// Extract IDs from []savedTrackInfo
	ids := make([]spotify.ID, 0, len(savedTracks))
	for _, t := range savedTracks {
		if t.ID != "" {
			ids = append(ids, t.ID)
		}
	}

	// Spotify API typically allows up to 100 IDs per request for audio features
	const batchSize = 100

	out := make(map[spotify.ID]*spotify.AudioFeatures, len(ids))

	for start := 0; start < len(ids); start += batchSize {
		end := start + batchSize
		if end > len(ids) {
			end = len(ids)
		}

		batch := ids[start:end]

		features, err := client.GetAudioFeatures(ctx, batch...)
		if err != nil {
			return nil, err
		}

		for i, f := range features {
			out[batch[i]] = f
		}
	}

	return out, nil
}
