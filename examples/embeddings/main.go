package main

import (
	"fmt"
	"log"

	"github.com/ZaguanLabs/mistral-go/v2/sdk"
)

func main() {
	client := sdk.NewMistralClientDefault("")

	// Simple embeddings
	fmt.Println("=== Simple Embeddings ===")
	texts := []string{
		"The quick brown fox jumps over the lazy dog",
		"A journey of a thousand miles begins with a single step",
	}

	response, err := client.Embeddings("mistral-embed", texts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated %d embeddings\n", len(response.Data))
	for i, embedding := range response.Data {
		fmt.Printf("Text %d: %d dimensions\n", i+1, len(embedding.Embedding))
	}

	// Embeddings with parameters
	fmt.Println("\n=== Embeddings with Parameters ===")
	encodingFormat := sdk.EncodingFormatFloat
	params := &sdk.EmbeddingRequest{
		Model:           "mistral-embed",
		Input:           texts,
		EncodingFormat:  &encodingFormat,
		OutputDimension: sdk.IntPtr(512), // Reduce dimensions
	}

	response, err = client.EmbeddingsWithParams("mistral-embed", texts, params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated %d embeddings with %d dimensions\n",
		len(response.Data),
		len(response.Data[0].Embedding))
	fmt.Printf("Total tokens used: %d\n", response.Usage.TotalTokens)
}
