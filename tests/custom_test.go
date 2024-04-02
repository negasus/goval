package tests

import (
	"testing"

	"github.com/negasus/goval"
)

func TestCustom(t *testing.T) {
	goval.AddCustomMessage(1, "en", "Custom error message {rule_value}")

	r := RequestCustom{
		ID: 0,
	}

	err := r.Validate()
	if len(err) != 1 {
		t.Fatalf("Expected 1 error, got %d", len(err))
	}

	fieldErr, ok := err["id"]
	if !ok {
		t.Fatalf("Key %s not found", "id")
	}

	if len(fieldErr) != 1 {
		t.Errorf("Expected 1 error for field %s, got %d", "id", len(fieldErr))
	}

	if fieldErr[0].Type != goval.ErrorTypeCustom {
		t.Errorf("Expected %s, got %s", goval.ErrorTypeCustom, fieldErr[0].Type)
	}

	if fieldErr[0].Values["rule_value"] != 1 {
		t.Errorf("Expected rule_value %v, got %d", 1, fieldErr[0].Values["rule_value"])
	}

	if fieldErr[0].String() != "Custom error message 1" {
		t.Errorf("Expected Custom error message 1, got %s", fieldErr[0].String())
	}
}

func TestCustomEmbed(t *testing.T) {
	r := RequestCustomEmbed{
		RequestCustomEmbedItem: RequestCustomEmbedItem{
			ID: 5,
		},
	}

	if err := testOneError(r.Validate(), "id", goval.ErrorTypeMinNumeric, 10); err != nil {
		t.Fatal(err)
	}

	r.ID = 15

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}
