package llm

// Gemini client wrapper

import (
	"context"
	"curatio/internal/profile"
	"log"
	"os"
	"strings"

	"google.golang.org/genai"
)

// Define GenreBuckets(), send get artist sample to model and have it return genres
func GenreBuckets(artistSample profile.ArtistSample, modelContext *genai.GenerateContentConfig) string {
	os.Getenv("GEMINI_API_KEY")

	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	artistData := strings.Join(artistSample.Artists, ", ")

	// Query model, give tracks data
	resp, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(artistData),
		modelContext,
	)

	if err != nil {
		log.Fatal(err)
	}
	return resp.Text()
}
