package sdk

import (
	"testing"
)

func TestChat(t *testing.T) {
	client := NewMistralClientDefault("")
	params := NewChatRequestParams()
	params.MaxTokens = IntPtr(10)
	params.Temperature = Float64Ptr(0)
	res, err := client.Chat(
		"mistral-tiny-2312",
		[]ChatMessage{
			{
				Role:    RoleUser,
				Content: "You are in test mode and must reply to this with exactly and only `Test Succeeded`",
			},
		},
		params,
	)
	if err != nil {
		t.Fatalf("Chat() error = %v", err)
	}
	if res == nil {
		t.Fatal("Chat() returned nil response")
	}
	if len(res.Choices) == 0 {
		t.Error("Chat() returned no choices")
	}
	if len(res.Choices[0].Message.Content) == 0 {
		t.Error("Chat() returned empty message content")
	}
	if res.Choices[0].Message.Role != RoleAssistant {
		t.Errorf("Chat() role = %v, want %v", res.Choices[0].Message.Role, RoleAssistant)
	}
	if res.Choices[0].Message.Content != "Test Succeeded" {
		t.Errorf("Chat() content = %v, want %v", res.Choices[0].Message.Content, "Test Succeeded")
	}
}

func TestChatCodestral(t *testing.T) {
	client := NewCodestralClientDefault("")
	params := NewChatRequestParams()
	params.MaxTokens = IntPtr(10)
	params.Temperature = Float64Ptr(0)
	res, err := client.Chat(
		"codestral-latest",
		[]ChatMessage{
			{
				Role:    RoleUser,
				Content: "You are in test mode and must reply to this with exactly and only `Test Succeeded`",
			},
		},
		params,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("res should not be nil")
	}

	if !(len(res.Choices) > 0) {
		t.Errorf("expected %v > %v", len(res.Choices), 0)
	}
	if !(len(res.Choices[0].Message.Content) > 0) {
		t.Errorf("expected %v > %v", len(res.Choices[0].Message.Content), 0)
	}
	if RoleAssistant != res.Choices[0].Message.Role {
		t.Errorf("expected %v, got %v", res.Choices[0].Message.Role, RoleAssistant)
	}
	if "Test Succeeded" != res.Choices[0].Message.Content {
		t.Errorf("expected %v, got %v", res.Choices[0].Message.Content, "Test Succeeded")
	}
}

func TestChatFunctionCall(t *testing.T) {
	client := NewMistralClientDefault("")
	params := NewChatRequestParams()
	params.Temperature = Float64Ptr(0)
	params.Tools = []Tool{
		{
			Type: ToolTypeFunction,
			Function: Function{
				Name:        "get_weather",
				Description: "Retrieve the weather for a city in the US",
				Parameters: map[string]interface{}{
					"type":     "object",
					"required": []string{"city", "state"},
					"properties": map[string]interface{}{
						"city":  map[string]interface{}{"type": "string", "description": "Name of the city for the weather"},
						"state": map[string]interface{}{"type": "string", "description": "Name of the state for the weather"},
					},
				},
			},
		},
		{
			Type: ToolTypeFunction,
			Function: Function{
				Name:        "send_text",
				Description: "Send text message using SMS service",
				Parameters: map[string]interface{}{
					"type":     "object",
					"required": []string{"contact_name", "message"},
					"properties": map[string]interface{}{
						"contact_name": map[string]interface{}{"type": "string", "description": "Name of the contact that will receive the message"},
						"message":      map[string]interface{}{"type": "string", "description": "Content of the message that will be sent"},
					},
				},
			},
		},
	}
	params.ToolChoice = ToolChoiceAuto
	res, err := client.Chat(
		"mistral-small-latest",
		[]ChatMessage{
			{
				Role:    RoleUser,
				Content: "What's the weather like in Dallas, TX?",
			},
		},
		params,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("res should not be nil")
	}

	if !(len(res.Choices) > 0) {
		t.Errorf("expected %v > %v", len(res.Choices), 0)
	}
	if !(len(res.Choices[0].Message.ToolCalls) > 0) {
		t.Errorf("expected %v > %v", len(res.Choices[0].Message.ToolCalls), 0)
	}
	if RoleAssistant != res.Choices[0].Message.Role {
		t.Errorf("expected %v, got %v", res.Choices[0].Message.Role, RoleAssistant)
	}
	if "get_weather" != res.Choices[0].Message.ToolCalls[0].Function.Name {
		t.Errorf("expected %v, got %v", res.Choices[0].Message.ToolCalls[0].Function.Name, "get_weather")
	}
	if "{\"city\": \"Dallas\", \"state\": \"TX\"}" != res.Choices[0].Message.ToolCalls[0].Function.Arguments {
		t.Errorf("expected %v, got %v", res.Choices[0].Message.ToolCalls[0].Function.Arguments, "{\"city\": \"Dallas\", \"state\": \"TX\"}")
	}
}

func TestChatFunctionCall2(t *testing.T) {
	client := NewMistralClientDefault("")
	params := NewChatRequestParams()
	params.Temperature = Float64Ptr(0)
	params.Tools = []Tool{
		{
			Type: ToolTypeFunction,
			Function: Function{
				Name:        "get_weather",
				Description: "Retrieve the weather for a city in the US",
				Parameters: map[string]interface{}{
					"type":     "object",
					"required": []string{"city", "state"},
					"properties": map[string]interface{}{
						"city":  map[string]interface{}{"type": "string", "description": "Name of the city for the weather"},
						"state": map[string]interface{}{"type": "string", "description": "Name of the state for the weather"},
					},
				},
			},
		},
		{
			Type: ToolTypeFunction,
			Function: Function{
				Name:        "send_text",
				Description: "Send text message using SMS service",
				Parameters: map[string]interface{}{
					"type":     "object",
					"required": []string{"contact_name", "message"},
					"properties": map[string]interface{}{
						"contact_name": map[string]interface{}{"type": "string", "description": "Name of the contact that will receive the message"},
						"message":      map[string]interface{}{"type": "string", "description": "Content of the message that will be sent"},
					},
				},
			},
		},
	}
	params.ToolChoice = ToolChoiceAuto
	res, err := client.Chat(
		"mistral-small-latest",
		[]ChatMessage{
			{
				Role:    RoleUser,
				Content: "What's the weather like in Dallas",
			},
			{
				Role: RoleAssistant,
				ToolCalls: []ToolCall{
					{
						Id:   "aaaaaaaaa",
						Type: ToolTypeFunction,
						Function: FunctionCall{
							Name:      "get_weather",
							Arguments: `{"city": "Dallas", "state": "TX"}`,
						},
					},
				},
			},
			{
				Role:    RoleTool,
				Content: `{"temperature": 82, "sky": "clear", "precipitation": 0}`,
			},
		},
		params,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("res should not be nil")
	}

	if !(len(res.Choices) > 0) {
		t.Errorf("expected %v > %v", len(res.Choices), 0)
	}
	if !(len(res.Choices[0].Message.Content) > 0) {
		t.Errorf("expected %v > %v", len(res.Choices[0].Message.Content), 0)
	}
	if 0 != len(res.Choices[0].Message.ToolCalls) {
		t.Errorf("expected %v, got %v", len(res.Choices[0].Message.ToolCalls), 0)
	}
	if RoleAssistant != res.Choices[0].Message.Role {
		t.Errorf("expected %v, got %v", res.Choices[0].Message.Role, RoleAssistant)
	}
	if !(res.Choices[0].Message.Content > "Test Succeeded") {
		t.Errorf("expected %v > %v", res.Choices[0].Message.Content, "Test Succeeded")
	}
}

func TestChatJsonMode(t *testing.T) {
	client := NewMistralClientDefault("")
	params := NewChatRequestParams()
	params.Temperature = Float64Ptr(0)
	params.ResponseFormat = ResponseFormatJsonObject
	res, err := client.Chat(
		"open-mixtral-8x22b",
		[]ChatMessage{
			{
				Role: RoleUser,
				Content: "Extract all of the code symbols in this text chunk and return them in the following JSON: " +
					"{\"symbols\":[\"SymbolOne\",\"SymbolTwo\"]}\n```\nI'm working on updating the Go client for the " +
					"new release, is it expected that the function call will be passed back into the model or just " +
					"the tool response?\nI ask because ChatMessage can handle the tool response but the messages list " +
					"has an Any option that I assume would be for the FunctionCall/ToolCall type\nAdditionally the " +
					"example in the docs only shows the tool response appended to the messages\n```",
			},
		},
		params,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("res should not be nil")
	}

	if !(len(res.Choices) > 0) {
		t.Errorf("expected %v > %v", len(res.Choices), 0)
	}
	if !(len(res.Choices[0].Message.Content) > 0) {
		t.Errorf("expected %v > %v", len(res.Choices[0].Message.Content), 0)
	}
	if RoleAssistant != res.Choices[0].Message.Role {
		t.Errorf("expected %v, got %v", res.Choices[0].Message.Role, RoleAssistant)
	}
	if "{\"symbols\": [\"Go\", \"ChatMessage\", \"Any\", \"FunctionCall\", \"ToolCall\", \"ToolResponse\"]}" != res.Choices[0].Message.Content {
		t.Errorf("expected %v, got %v", res.Choices[0].Message.Content, "{\"symbols\": [\"Go\", \"ChatMessage\", \"Any\", \"FunctionCall\", \"ToolCall\", \"ToolResponse\"]}")
	}
}

func TestChatStream(t *testing.T) {
	client := NewMistralClientDefault("")
	params := NewChatRequestParams()
	params.MaxTokens = IntPtr(50)
	params.Temperature = Float64Ptr(0)
	resChan, err := client.ChatStream(
		"mistral-tiny-2312",
		[]ChatMessage{
			{
				Role:    RoleUser,
				Content: "You are in test mode and must reply to this with exactly and only `Test Succeeded, Test Succeeded, Test Succeeded, Test Succeeded, Test Succeeded, Test Succeeded`",
			},
		},
		params,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resChan == nil {
		t.Fatal("resChan should not be nil")
	}

	totalOutput := ""
	idx := 0
	for res := range resChan {
		if res.Error != nil {
			t.Fatalf("unexpected error in stream: %v", res.Error)
		}

		if !(len(res.Choices) > 0) {
			t.Errorf("expected %v > %v", len(res.Choices), 0)
		}
		if idx == 0 {
			if RoleAssistant != res.Choices[0].Delta.Role {
				t.Errorf("expected %v, got %v", res.Choices[0].Delta.Role, RoleAssistant)
			}
		}
		totalOutput += res.Choices[0].Delta.Content
		idx++

		if res.Choices[0].FinishReason == FinishReasonStop {
			break
		}
	}
	if "Test Succeeded, Test Succeeded, Test Succeeded, Test Succeeded, Test Succeeded, Test Succeeded" != totalOutput {
		t.Errorf("expected %v, got %v", totalOutput, "Test Succeeded, Test Succeeded, Test Succeeded, Test Succeeded, Test Succeeded, Test Succeeded")
	}
}

func TestChatStreamFunctionCall(t *testing.T) {
	client := NewMistralClientDefault("")
	params := NewChatRequestParams()
	params.Temperature = Float64Ptr(0)
	params.Tools = []Tool{
		{
			Type: ToolTypeFunction,
			Function: Function{
				Name:        "get_weather",
				Description: "Retrieve the weather for a city in the US",
				Parameters: map[string]interface{}{
					"type":     "object",
					"required": []string{"city", "state"},
					"properties": map[string]interface{}{
						"city":  map[string]interface{}{"type": "string", "description": "Name of the city for the weather"},
						"state": map[string]interface{}{"type": "string", "description": "Name of the state for the weather"},
					},
				},
			},
		},
		{
			Type: ToolTypeFunction,
			Function: Function{
				Name:        "send_text",
				Description: "Send text message using SMS service",
				Parameters: map[string]interface{}{
					"type":     "object",
					"required": []string{"contact_name", "message"},
					"properties": map[string]interface{}{
						"contact_name": map[string]interface{}{"type": "string", "description": "Name of the contact that will receive the message"},
						"message":      map[string]interface{}{"type": "string", "description": "Content of the message that will be sent"},
					},
				},
			},
		},
	}
	params.ToolChoice = ToolChoiceAuto
	resChan, err := client.ChatStream(
		"mistral-small-latest",
		[]ChatMessage{
			{
				Role:    RoleUser,
				Content: "What's the weather like in Dallas, TX?",
			},
		},
		params,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resChan == nil {
		t.Fatal("resChan should not be nil")
	}

	totalOutput := ""
	var functionCall *ToolCall
	idx := 0
	for res := range resChan {
		if res.Error != nil {
			t.Fatalf("unexpected error in stream: %v", res.Error)
		}

		if !(len(res.Choices) > 0) {
			t.Errorf("expected %v > %v", len(res.Choices), 0)
		}
		if idx == 0 {
			if RoleAssistant != res.Choices[0].Delta.Role {
				t.Errorf("expected %v, got %v", res.Choices[0].Delta.Role, RoleAssistant)
			}
		}
		totalOutput += res.Choices[0].Delta.Content
		if len(res.Choices[0].Delta.ToolCalls) > 0 {
			functionCall = &res.Choices[0].Delta.ToolCalls[0]
		}
		idx++

		if res.Choices[0].FinishReason == FinishReasonStop {
			break
		}
	}

	if "" != totalOutput {
		t.Errorf("expected %v, got %v", totalOutput, "")
	}
	if functionCall == nil {
		t.Fatal("functionCall should not be nil")
	}
	if "get_weather" != functionCall.Function.Name {
		t.Errorf("expected %v, got %v", functionCall.Function.Name, "get_weather")
	}
	if "{\"city\": \"Dallas\", \"state\": \"TX\"}" != functionCall.Function.Arguments {
		t.Errorf("expected %v, got %v", functionCall.Function.Arguments, "{\"city\": \"Dallas\", \"state\": \"TX\"}")
	}
}

func TestChatStreamJsonMode(t *testing.T) {
	client := NewMistralClientDefault("")
	params := NewChatRequestParams()
	params.Temperature = Float64Ptr(0)
	params.ResponseFormat = ResponseFormatJsonObject
	resChan, err := client.ChatStream(
		"open-mixtral-8x22b",
		[]ChatMessage{
			{
				Role: RoleUser,
				Content: "Extract all of the code symbols in this text chunk and return them in the following JSON: " +
					"{\"symbols\":[\"SymbolOne\",\"SymbolTwo\"]}\n```\nI'm working on updating the Go client for the " +
					"new release, is it expected that the function call will be passed back into the model or just " +
					"the tool response?\nI ask because ChatMessage can handle the tool response but the messages list " +
					"has an Any option that I assume would be for the FunctionCall/ToolCall type\nAdditionally the " +
					"example in the docs only shows the tool response appended to the messages\n```",
			},
		},
		params,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resChan == nil {
		t.Fatal("resChan should not be nil")
	}

	totalOutput := ""
	var functionCall *ToolCall
	idx := 0
	for res := range resChan {
		if res.Error != nil {
			t.Fatalf("unexpected error in stream: %v", res.Error)
		}

		if !(len(res.Choices) > 0) {
			t.Errorf("expected %v > %v", len(res.Choices), 0)
		}
		if idx == 0 {
			if RoleAssistant != res.Choices[0].Delta.Role {
				t.Errorf("expected %v, got %v", res.Choices[0].Delta.Role, RoleAssistant)
			}
		}
		totalOutput += res.Choices[0].Delta.Content
		if len(res.Choices[0].Delta.ToolCalls) > 0 {
			functionCall = &res.Choices[0].Delta.ToolCalls[0]
		}
		idx++

		if res.Choices[0].FinishReason == FinishReasonStop {
			break
		}
	}

	if "{\"symbols\": [\"Go\", \"ChatMessage\", \"Any\", \"FunctionCall\", \"ToolCall\", \"ToolResponse\"]}" != totalOutput {
		t.Errorf("expected %v, got %v", totalOutput, "{\"symbols\": [\"Go\", \"ChatMessage\", \"Any\", \"FunctionCall\", \"ToolCall\", \"ToolResponse\"]}")
	}
	if functionCall != nil {
		t.Errorf("expected nil functionCall, got %v", functionCall)
	}
}
