package tests

import (
	"fmt"
	"testing"

	"github.com/negasus/goval"
)

func TestMinInt(t *testing.T) {
	r := RequestInt{
		ID: 0,
	}

	if err := testOneError(r.Validate(), "id", goval.ErrorTypeMinNumeric, 1); err != nil {
		t.Fatal(err)
	}

	r.ID = 1

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func TestMinFloat(t *testing.T) {
	r := RequestFloat{
		Price: 0.5,
	}

	if err := testOneError(r.Validate(), "price", goval.ErrorTypeMinNumeric, 10.5); err != nil {
		t.Fatal(err)
	}

	r.Price = 11.1

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func TestMinString(t *testing.T) {
	r := RequestString{
		Name: "a",
	}

	if err := testOneError(r.Validate(), "name", goval.ErrorTypeMinString, 3); err != nil {
		t.Fatal(err)
	}

	r.Name = "123456789"

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func TestMinStringArray(t *testing.T) {
	r := RequestStringArray{
		Keys: []string{"a"},
	}

	if err := testOneError(r.Validate(), "keys", goval.ErrorTypeMinArray, 3); err != nil {
		t.Fatal(err)
	}

	r.Keys = []string{"1", "2", "3", "4"}

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func TestMaxInt(t *testing.T) {
	r := RequestInt{
		ID: 12,
	}

	if err := testOneError(r.Validate(), "id", goval.ErrorTypeMaxNumeric, 10); err != nil {
		t.Fatal(err)
	}

	r.ID = 10

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func TestMaxFloat(t *testing.T) {
	r := RequestFloat{
		Price: 21.5,
	}

	if err := testOneError(r.Validate(), "price", goval.ErrorTypeMaxNumeric, 20.5); err != nil {
		t.Fatal(err)
	}

	r.Price = 20.1

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func TestMaxString(t *testing.T) {
	r := RequestString{
		Name: "12345678901",
	}

	if err := testOneError(r.Validate(), "name", goval.ErrorTypeMaxString, 10); err != nil {
		t.Fatal(err)
	}

	r.Name = "12345"

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func TestMaxStringArray(t *testing.T) {
	r := RequestStringArray{
		Keys: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
	}

	if err := testOneError(r.Validate(), "keys", goval.ErrorTypeMaxArray, 10); err != nil {
		t.Fatal(err)
	}

	r.Keys = []string{"1", "2", "3", "4"}

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func TestMaxIntSliceMin(t *testing.T) {
	r := RequestIntSlice{
		IDs: []int{},
	}

	if err := testOneError(r.Validate(), "ids", goval.ErrorTypeMinArray, 1); err != nil {
		t.Fatal(err)
	}

	r.IDs = []int{0, 1, 2}

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func TestMaxIntSliceMax(t *testing.T) {
	r := RequestIntSlice{
		IDs: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
	}

	if err := testOneError(r.Validate(), "ids", goval.ErrorTypeMaxArray, 10); err != nil {
		t.Fatal(err)
	}

	r.IDs = []int{0, 1, 2}

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func testOneError(err goval.Errors, key string, errType goval.ErrorType, ruleValue any) error {
	if len(err) != 1 {
		return fmt.Errorf("expected 1 error, got %d", len(err))
	}

	fieldErr, ok := err[key]
	if !ok {
		return fmt.Errorf("key %s not found", key)
	}

	if len(fieldErr) != 1 {
		return fmt.Errorf("expected 1 error for field %s, got %d", key, len(fieldErr))
	}

	if fieldErr[0].Type != errType {
		return fmt.Errorf("expected %s, got %s", errType, fieldErr[0].Type)
	}

	if ruleValue != nil && fieldErr[0].Values["rule_value"] != ruleValue {
		return fmt.Errorf("expected rule_value %v, got %d", ruleValue, fieldErr[0].Values["rule_value"])
	}

	return nil
}
