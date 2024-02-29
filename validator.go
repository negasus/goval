package goval

import "fmt"

type Validator struct {
}

func New() *Validator {
	v := &Validator{}
	return v
}

func (v *Validator) Generate() error {
	fmt.Printf("validator generate\n")
	return nil
}
