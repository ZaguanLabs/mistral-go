package sdk

import (
	"testing"
)

func TestListModels(t *testing.T) {
	client := NewMistralClientDefault("")

	_, err := client.ListModels()

	if err == nil {
		t.Log("ListModels succeeded")
	} else {
		t.Logf("ListModels failed as expected: %v", err)
	}
}

func TestModelPermissionStructure(t *testing.T) {
	perm := ModelPermission{
		ID:                 "perm-123",
		Object:             "model_permission",
		Created:            1234567890,
		AllowCreateEngine:  true,
		AllowSampling:      true,
		AllowLogprobs:      false,
		AllowSearchIndices: true,
		AllowView:          true,
		AllowFineTuning:    false,
		Organization:       "org-123",
		Group:              "group-123",
		IsBlocking:         false,
	}

	if !perm.AllowCreateEngine {
		t.Error("AllowCreateEngine should be true")
	}

	if perm.AllowLogprobs {
		t.Error("AllowLogprobs should be false")
	}
}

func TestModelCardStructure(t *testing.T) {
	card := ModelCard{
		ID:      "model-123",
		Object:  "model",
		Created: 1234567890,
		OwnedBy: "mistral",
		Root:    "root-model",
		Parent:  "parent-model",
		Permission: []ModelPermission{
			{ID: "perm-1", AllowView: true},
			{ID: "perm-2", AllowSampling: true},
		},
	}

	if card.OwnedBy != "mistral" {
		t.Error("OwnedBy not set correctly")
	}

	if len(card.Permission) != 2 {
		t.Error("Permission list length incorrect")
	}
}

func TestModelListStructure(t *testing.T) {
	list := ModelList{
		Object: "list",
		Data: []ModelCard{
			{ID: "model-1", OwnedBy: "mistral"},
			{ID: "model-2", OwnedBy: "mistral"},
			{ID: "model-3", OwnedBy: "community"},
		},
	}

	if list.Object != "list" {
		t.Error("Object type should be 'list'")
	}

	if len(list.Data) != 3 {
		t.Error("Data length should be 3")
	}
}

func TestModelCardOptionalFields(t *testing.T) {
	// Test ModelCard with minimal fields
	card := ModelCard{
		ID:         "model-123",
		Object:     "model",
		Created:    123,
		OwnedBy:    "mistral",
		Permission: []ModelPermission{},
	}

	if card.Root != "" {
		t.Error("Root should be empty when not set")
	}

	if card.Parent != "" {
		t.Error("Parent should be empty when not set")
	}
}
