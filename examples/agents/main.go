package main

import (
	"fmt"
	"log"

	"github.com/ZaguanLabs/mistral-go/v2/sdk"
)

func main() {
	client := sdk.NewMistralClientDefault("")

	// Create a Mistral Agent
	fmt.Println("=== Creating Mistral Agent ===")
	agent, err := client.CreateMistralAgent(&sdk.CreateMistralAgentRequest{
		Model:        "mistral-large-latest",
		Name:         sdk.StringPtr("Code Review Assistant"),
		Description:  sdk.StringPtr("An agent that helps review code"),
		Instructions: sdk.StringPtr("You are an expert code reviewer. Provide constructive feedback on code quality, best practices, and potential bugs."),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created agent: %s (ID: %s)\n\n", *agent.Name, agent.ID)

	// List all agents
	fmt.Println("=== Listing Agents ===")
	agents, err := client.ListMistralAgents(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d agents\n", len(agents.Data))
	for _, a := range agents.Data {
		if a.Name != nil {
			fmt.Printf("- %s (ID: %s)\n", *a.Name, a.ID)
		}
	}

	// Use agent for completion
	fmt.Println("\n=== Agent Completion ===")
	messages := []sdk.ChatMessage{
		sdk.UserMessage("Review this code: def add(a, b): return a + b"),
	}

	response, err := client.AgentComplete(agent.ID, messages, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Agent response: %s\n", response.Choices[0].Message.Content)
}
