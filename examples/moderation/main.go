package main

import (
	"fmt"
	"log"

	"github.com/ZaguanLabs/mistral-go/v2/sdk"
)

func main() {
	client := sdk.NewMistralClientDefault("")

	// Content moderation
	fmt.Println("=== Content Moderation ===")
	texts := []string{
		"This is a perfectly normal message.",
		"Let's discuss the weather today.",
	}

	response, err := client.ModerateText("mistral-moderation-latest", texts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Moderated %d texts\n", len(response.Results))
	for i, result := range response.Results {
		fmt.Printf("\nText %d:\n", i+1)
		for _, category := range result.Categories {
			fmt.Printf("  - %s: %.4f\n", category.CategoryName, category.Score)
		}
	}
}
