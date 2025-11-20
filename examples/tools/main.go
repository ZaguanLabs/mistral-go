package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ZaguanLabs/mistral-go/v2/sdk"
)

func main() {
	client := sdk.NewMistralClientDefault("")

	// Define a weather tool
	params := sdk.NewChatRequestParams()
	params.Tools = []sdk.Tool{
		{
			Type: "function",
			Function: sdk.Function{
				Name:        "get_weather",
				Description: "Get the current weather in a location",
				Parameters: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"location": map[string]interface{}{
							"type":        "string",
							"description": "The city and state, e.g. San Francisco, CA",
						},
					},
					"required": []string{"location"},
				},
			},
		},
	}
	params.ToolChoice = "auto"

	fmt.Println("=== Tool Calling Example ===")
	response, err := client.Chat(
		"mistral-large-latest",
		[]sdk.ChatMessage{
			sdk.UserMessage("What's the weather like in Paris?"),
		},
		params,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the model wants to call a tool
	if len(response.Choices) > 0 && len(response.Choices[0].Message.ToolCalls) > 0 {
		toolCall := response.Choices[0].Message.ToolCalls[0]
		fmt.Printf("Tool called: %s\n", toolCall.Function.Name)
		fmt.Printf("Arguments: %s\n", toolCall.Function.Arguments)

		// Parse arguments
		var args map[string]interface{}
		if err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Location: %s\n", args["location"])

		// Simulate tool response
		toolResponse := `{"temperature": 22, "condition": "sunny"}`

		// Send tool response back to model
		messages := []sdk.ChatMessage{
			sdk.UserMessage("What's the weather like in Paris?"),
			response.Choices[0].Message,
			sdk.ToolMessage(toolCall.Id, toolResponse),
		}

		finalResponse, err := client.Chat("mistral-large-latest", messages, params)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\nFinal Response: %s\n", finalResponse.Choices[0].Message.Content)
	}
}
