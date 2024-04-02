package tests

import (
	"testing"

	"github.com/negasus/goval"
)

func TestInSliceIntVar(t *testing.T) {
	r := CompareInIntVar{
		ID: 42,
	}

	if err := testOneError(r.Validate(), "id", goval.ErrorTypeIn, nil); err != nil {
		t.Fatal(err)
	}

	r.ID = 2

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func TestInSliceIntVarSlice(t *testing.T) {
	r := CompareInIntVarSlice{
		ID: []int{54},
	}

	if err := testOneError(r.Validate(), "id", goval.ErrorTypeIn, nil); err != nil {
		t.Fatal(err)
	}

	r.ID = []int{3}

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func TestInSliceInt(t *testing.T) {
	r := CompareInInt{
		ID: 42,
	}

	if err := testOneError(r.Validate(), "id", goval.ErrorTypeIn, nil); err != nil {
		t.Fatal(err)
	}

	r.ID = 20

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func TestInSliceIntSlice(t *testing.T) {
	r := CompareInIntSlice{
		ID: []int{42},
	}

	if err := testOneError(r.Validate(), "id", goval.ErrorTypeIn, nil); err != nil {
		t.Fatal(err)
	}

	r.ID = []int{20}

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func TestInSliceStringVar(t *testing.T) {
	r := CompareInStringVar{
		Name: "x",
	}

	if err := testOneError(r.Validate(), "name", goval.ErrorTypeIn, nil); err != nil {
		t.Fatal(err)
	}

	r.Name = "a"

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func TestInSliceStringVarSlice(t *testing.T) {
	r := CompareInStringVarSlice{
		Name: []string{"x"},
	}

	if err := testOneError(r.Validate(), "name", goval.ErrorTypeIn, nil); err != nil {
		t.Fatal(err)
	}

	r.Name = []string{"a"}

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func TestInSliceString(t *testing.T) {
	r := CompareInString{
		Name: "xx",
	}

	if err := testOneError(r.Validate(), "name", goval.ErrorTypeIn, nil); err != nil {
		t.Fatal(err)
	}

	r.Name = "aa"

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}

func TestInSliceStringSlice(t *testing.T) {
	r := CompareInStringSlice{
		Name: []string{"xx"},
	}

	if err := testOneError(r.Validate(), "name", goval.ErrorTypeIn, nil); err != nil {
		t.Fatal(err)
	}

	r.Name = []string{"aa"}

	err := r.Validate()
	if len(err) != 0 {
		t.Errorf("Expected 0 error, got %d", len(err))
	}
}
