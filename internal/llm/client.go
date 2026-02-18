package llm

// Gemini client wrapper

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/genai"
)

func gemini_api_call() {
	os.Getenv("GEMINI_API_KEY")

	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Ask User for a message
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Whaddup? ...")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)

	// Query model
	resp, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(userInput),
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Text())
}
