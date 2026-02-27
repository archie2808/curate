package profile

import (
	"curatio/internal/spotify"
)

type ArtistSample struct {
	Artists []string
}
// SavedTrackSample

// Get the first 500 tracks from []SavedTrackInfo, this should be provide enough diversity for llm to define representative genre buckets
func GetArtistSample(trackInfo []spotify.SavedTrackInfo) ArtistSample {
	artistNamesNoDupes := make(map[string]struct{})
	if len(trackInfo) > 500 {

		for _, tracks := range trackInfo[:500] {

			for _, artists := range tracks.Artists {
				artistNamesNoDupes[artists.Name] = struct{}{}
			}
		}

	} else {

		for _, tracks := range trackInfo[:len(trackInfo)/2] {

			for _, artists := range tracks.Artists {
				artistNamesNoDupes[artists.Name] = struct{}{}
			}

		}
	}

	artistNames := make([]string, 0, len(artistNamesNoDupes))
	for key := range artistNamesNoDupes {
		artistNames = append(artistNames, key)
	}
	
  	return ArtistSample{Artists: artistNames}

}
