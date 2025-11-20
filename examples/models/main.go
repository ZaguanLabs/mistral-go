package main

import (
	"fmt"
	"log"

	"github.com/ZaguanLabs/mistral-go/v2/sdk"
)

func main() {
	client := sdk.NewMistralClientDefault("")

	// List all models
	fmt.Println("=== Available Models ===")
	models, err := client.ListModels()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d models:\n\n", len(models.Data))
	for _, model := range models.Data {
		fmt.Printf("ID: %s\n", model.ID)
		fmt.Printf("  Object: %s\n", model.Object)
		fmt.Printf("  Owned by: %s\n", model.OwnedBy)
		fmt.Printf("  Created: %d\n", model.Created)
		if model.Root != "" {
			fmt.Printf("  Root: %s\n", model.Root)
		}
		fmt.Println()
	}

	// List models is the primary API - individual model details
	// are included in the list response
	fmt.Println("=== Model Details ===")
	if len(models.Data) > 0 {
		firstModel := models.Data[0]
		fmt.Printf("First model ID: %s\n", firstModel.ID)
		fmt.Printf("Permissions: %d\n", len(firstModel.Permission))
	}
}
