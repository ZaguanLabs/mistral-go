package main

import (
	"fmt"
	"log"

	"github.com/ZaguanLabs/mistral-go/v2/sdk"
)

func main() {
	client := sdk.NewCodestralClientDefault("")

	// Simple FIM completion
	fmt.Println("=== Fill-in-the-Middle Code Completion ===")
	response, err := client.FIM(&sdk.FIMRequestParams{
		Model:  "codestral-latest",
		Prompt: "def fibonacci(n):",
		Suffix: sdk.StringPtr("    return result"),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Completion: %s\n\n", response.Choices[0].Message.Content)

	// FIM with parameters
	fmt.Println("=== FIM with Parameters ===")
	response, err = client.FIM(&sdk.FIMRequestParams{
		Model:       "codestral-latest",
		Prompt:      "function calculateSum(a, b) {",
		Suffix:      sdk.StringPtr("}"),
		Temperature: sdk.Float64Ptr(0.3),
		MaxTokens:   sdk.IntPtr(50),
		Stop:        []string{"\n\n"},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Completion: %s\n\n", response.Choices[0].Message.Content)

	// Streaming FIM
	fmt.Println("=== Streaming FIM ===")
	stream, err := client.FIMStream(&sdk.FIMRequestParams{
		Model:  "codestral-latest",
		Prompt: "class Calculator:",
		Suffix: sdk.StringPtr("    pass"),
	})
	if err != nil {
		log.Fatal(err)
	}

	for chunk := range stream {
		if chunk.Error != nil {
			log.Fatal(chunk.Error)
		}
		if len(chunk.Choices) > 0 {
			fmt.Print(chunk.Choices[0].Delta.Content)
		}
	}
	fmt.Println()
}
