package tests

import (
	"testing"

	"github.com/negasus/goval"
)

func TestMinInt(t *testing.T) {
	r := RequestInt{
		ID: 0,
	}

	if !testOneError(t, r.Validate(), "id", goval.ErrorTypeMinNumeric, 1) {
		return
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

	if !testOneError(t, r.Validate(), "price", goval.ErrorTypeMinNumeric, 10.5) {
		return
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

	if !testOneError(t, r.Validate(), "name", goval.ErrorTypeMinString, 3) {
		return
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

	if !testOneError(t, r.Validate(), "keys", goval.ErrorTypeMinArray, 3) {
		return
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

	if !testOneError(t, r.Validate(), "id", goval.ErrorTypeMaxNumeric, 10) {
		return
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

	if !testOneError(t, r.Validate(), "price", goval.ErrorTypeMaxNumeric, 20.5) {
		return
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

	if !testOneError(t, r.Validate(), "name", goval.ErrorTypeMaxString, 10) {
		return
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

	if !testOneError(t, r.Validate(), "keys", goval.ErrorTypeMaxArray, 10) {
		return
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

	if !testOneError(t, r.Validate(), "ids", goval.ErrorTypeMinArray, 1) {
		return
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

	if !testOneError(t, r.Validate(), "ids", goval.ErrorTypeMaxArray, 10) {
		return
	}

	r.IDs = []int{0, 1, 2}

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func testOneError(t *testing.T, err goval.Errors, key string, errType goval.ErrorType, ruleValue any) bool {
	if len(err) != 1 {
		t.Errorf("Expected 1 error, got %d", len(err))
		return false
	}

	fieldErr, ok := err[key]
	if !ok {
		t.Errorf("Key %s not found", key)
		return false
	}

	if len(fieldErr) != 1 {
		t.Errorf("Expected 1 error for field %s, got %d", key, len(fieldErr))
		return false
	}

	if fieldErr[0].Type != errType {
		t.Errorf("Expected %s, got %s", errType, fieldErr[0].Type)
		return false
	}

	if ruleValue != nil && fieldErr[0].Values["rule_value"] != ruleValue {
		t.Errorf("Expected rule_value %v, got %d", ruleValue, fieldErr[0].Values["rule_value"])
		return false
	}

	return true
}
