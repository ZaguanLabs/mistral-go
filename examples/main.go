package main

import (
	"log"

	sdk "github.com/ZaguanLabs/mistral-go/v2/sdk"
)

func main() {
	// If api key is empty it will load from MISTRAL_API_KEY env var
	client := sdk.NewMistralClientDefault("your-api-key")

	// Example: Using Chat Completions
	chatRes, err := client.Chat("mistral-small-latest", []sdk.ChatMessage{{Content: "Hello, world!", Role: sdk.RoleUser}}, nil)
	if err != nil {
		log.Fatalf("Error getting chat completion: %v", err)
	}
	log.Printf("Chat completion: %+v\n", chatRes)

	// Example: Using Chat Completions Stream
	chatResChan, err := client.ChatStream("mistral-small-latest", []sdk.ChatMessage{{Content: "Hello, world!", Role: sdk.RoleUser}}, nil)
	if err != nil {
		log.Fatalf("Error getting chat completion stream: %v", err)
	}

	for chatResChunk := range chatResChan {
		if chatResChunk.Error != nil {
			log.Fatalf("Error while streaming response: %v", chatResChunk.Error)
		}
		log.Printf("Chat completion stream part: %+v\n", chatResChunk)
	}

	// Example: Using Embeddings
	embsRes, err := client.Embeddings("mistral-embed", []string{"Embed this sentence.", "As well as this one."})
	if err != nil {
		log.Fatalf("Error getting embeddings: %v", err)
	}

	log.Printf("Embeddings response: %+v\n", embsRes)
}
