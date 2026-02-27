package llm

import "github.com/zmb3/spotify/v2"

// PlaylistPlan, PlaylistSpec, TrackSuggestion structs + genai schema

type PlaylistPlan struct {
	Artist string
}

type PlaylistSpec struct {
}

type TrackSuggestion struct {
	trackID spotify.ID
}
