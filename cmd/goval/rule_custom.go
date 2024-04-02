package main

import (
	"strings"
)

func getCustomFunc(structFieldName, fieldName, rule string) (string, error) {
	funcName := strings.TrimPrefix(rule, "@")

	if funcName == "" {
		funcName = structFieldName + ".Validate"
	}

	data := map[string]any{
		"field":    structFieldName,
		"name":     fieldName,
		"funcName": funcName,
	}

	return returnWithTemplate("call_custom", data)
}

func getCustomArrayFunc(structFieldName, fieldName, rule string) (string, error) {
	rd.Imports["fmt"] = struct{}{}

	funcName := strings.TrimPrefix(rule, "@")

	if funcName == "" {
		funcName = "Validate"
	}

	data := map[string]any{
		"field":    structFieldName,
		"name":     fieldName,
		"funcName": funcName,
	}

	return returnWithTemplate("call_custom_array", data)
}
