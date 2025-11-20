package main

import (
	"fmt"
	"log"

	"github.com/ZaguanLabs/mistral-go/v2/sdk"
)

func main() {
	// Initialize client (loads from MISTRAL_API_KEY env var)
	client := sdk.NewMistralClientDefault("")

	// Simple chat completion
	fmt.Println("=== Simple Chat ===")
	response, err := client.Chat(
		"mistral-small-latest",
		[]sdk.ChatMessage{
			sdk.UserMessage("What is the capital of France?"),
		},
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %s\n\n", response.Choices[0].Message.Content)

	// Chat with parameters
	fmt.Println("=== Chat with Parameters ===")
	params := sdk.NewChatRequestParams()
	params.Temperature = sdk.Float64Ptr(0.7)
	params.MaxTokens = sdk.IntPtr(100)

	response, err = client.Chat(
		"mistral-small-latest",
		[]sdk.ChatMessage{
			sdk.SystemMessage("You are a helpful assistant."),
			sdk.UserMessage("Tell me a short joke."),
		},
		params,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %s\n\n", response.Choices[0].Message.Content)

	// Streaming chat
	fmt.Println("=== Streaming Chat ===")
	streamParams := sdk.NewChatRequestParams()
	streamParams.Temperature = sdk.Float64Ptr(0.7)

	stream, err := client.ChatStream(
		"mistral-small-latest",
		[]sdk.ChatMessage{
			sdk.UserMessage("Count from 1 to 5."),
		},
		streamParams,
	)
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
