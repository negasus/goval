package string_in

import (
	"fmt"
	"testing"
)

func TestStringInVals(t *testing.T) {
	req := Request{}
	req.ID = -1

	err := req.Validate()

	fmt.Printf("%v\n", err)
}
