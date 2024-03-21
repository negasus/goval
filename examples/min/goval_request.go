// Code generated by goval. DO NOT EDIT.
// Used flags: -t Request

package main

import "github.com/negasus/goval"

func (model *Request) Validate() goval.Errors {
	errors := goval.Errors{}

	{ // Name
		if len(model.Name) < 3 {
			errors["Name"] = append(errors["Name"], goval.Error{
				Type: goval.ErrorTypeMinString,
				Values: map[string]any{
					"field": "Name",
					"min":   3,
				},
			})
		}
	}
	{ // Name
		if len(model.Name) > 10 {
			errors["Name"] = append(errors["Name"], goval.Error{
				Type: goval.ErrorTypeMaxString,
				Values: map[string]any{
					"field": "Name",
					"min":   10,
				},
			})
		}
	}
	{ // Age
		if model.Age < 18 {
			errors["Age"] = append(errors["Age"], goval.Error{
				Type: goval.ErrorTypeMinNumeric,
				Values: map[string]any{
					"field": "Age",
					"min":   18,
				},
			})
		}
	}

	return errors
}
