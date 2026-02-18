package spotify

import (
	"context"

	"github.com/zmb3/spotify/v2"
)

// FetchLikedSongs && FetchArtistGenres

type SavedTrackInfo struct {
	ID         spotify.ID             `json:"id"`
	Name       string                 `json:"name"`
	Artists    []spotify.SimpleArtist `json:"artists"`
	DurationMs int                    `json:"duration_ms"`
	Popularity int                    `json:"popularity"`
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

func FetchArtistGenres() {
	panic("Not implemented")
}
