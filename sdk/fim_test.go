package sdk

import (
	"testing"
)

func TestFIM(t *testing.T) {
	client := NewMistralClientDefault("")
	params := FIMRequestParams{
		Model:       "codestral-latest",
		Prompt:      "def f(",
		Suffix:      StringPtr("return a + b"),
		MaxTokens:   IntPtr(64),
		Temperature: Float64Ptr(0),
		Stop:        []string{"\n"},
	}
	res, err := client.FIM(&params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("res should not be nil")
	}

	if !(len(res.Choices) > 0) {
		t.Errorf("expected %v > %v", len(res.Choices), 0)
	}
	if res.Choices[0].Message.Content != "a, b):" {
		t.Errorf("expected %v, got %v", "a, b):", res.Choices[0].Message.Content)
	}
	if res.Choices[0].FinishReason != FinishReasonStop {
		t.Errorf("expected %v, got %v", FinishReasonStop, res.Choices[0].FinishReason)
	}
}

func TestFIMWithStop(t *testing.T) {
	client := NewMistralClientDefault("")
	params := FIMRequestParams{
		Model:       "codestral-latest",
		Prompt:      "def is_odd(n): \n return n % 2 == 1 \n def test_is_odd():",
		Suffix:      StringPtr("test_is_odd()"),
		MaxTokens:   IntPtr(64),
		Temperature: Float64Ptr(0),
		Stop:        []string{"False"},
	}
	res, err := client.FIM(&params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("res should not be nil")
	}

	if !(len(res.Choices) > 0) {
		t.Errorf("expected %v > %v", len(res.Choices), 0)
	}
	expected := "\n assert is_odd(1) == True\n assert is_odd(2) == "
	if res.Choices[0].Message.Content != expected {
		t.Errorf("expected %v, got %v", expected, res.Choices[0].Message.Content)
	}
	if res.Choices[0].FinishReason != FinishReasonStop {
		t.Errorf("expected %v, got %v", FinishReasonStop, res.Choices[0].FinishReason)
	}
}

func TestFIMInvalidModel(t *testing.T) {
	client := NewMistralClientDefault("")
	params := FIMRequestParams{
		Model:       "invalid-model",
		Prompt:      "This is a test prompt",
		Suffix:      StringPtr("This is a test suffix"),
		MaxTokens:   IntPtr(10),
		Temperature: Float64Ptr(0.5),
	}
	res, err := client.FIM(&params)
	if err == nil {
		t.Error("expected error for invalid model")
	}
	if res != nil {
		t.Error("expected nil response for invalid model")
	}
}
