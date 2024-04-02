package tests

import "testing"

func TestCompositeNames(t *testing.T) {
	r := &Composite{
		Item: Composite2{
			ID: 0,
		},
		Items: []Composite2{
			{ID: -1},
			{ID: 2},
			{ID: 0},
		},
	}
	_ = r
}
