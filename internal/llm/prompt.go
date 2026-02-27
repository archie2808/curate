package llm

import (
	"os"

	"google.golang.org/genai"
)

// System Prompt Instructions

func PromptInstructions() *genai.GenerateContentConfig {
	os.Getenv("GEMINI_API_KEY")

	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText("You are a curatio, a musical playlist curator, generate a max of 10 genres based on input data", ""),
	}
	return config

}
