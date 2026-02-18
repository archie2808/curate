# Curatio

A tool for curating Spotify playlists using AI.

## How it works

1. **Connects to your Spotify account** via OAuth and reads your liked songs library
2. **Builds a library profile** — genre distribution, top artists, era breakdown, and representative tracks — compressed into a compact summary
3. **You describe what you want** in plain English (e.g. "make me 3 gym playlists, 45 mins each, high energy, no slow songs")
4. **AI curates a playlist plan** using its own music knowledge and your library profile — suggesting specific tracks by name and artist
5. **Tracks are searched on Spotify**, matched, trimmed to target duration, and **playlists are created in your account**
6. **Iterate** — refine with follow-up requests ("more like this", "less pop", "add some 90s hip hop")

