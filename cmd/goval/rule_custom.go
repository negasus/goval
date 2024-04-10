package main

import (
	"strings"
)

func getCustomFunc(a ruleFuncArgs) (string, error) {
	funcName := strings.TrimPrefix(a.rule, "@")

	if funcName == "" {
		funcName = a.structFieldName + ".Validate"
	}

	data := map[string]any{
		"field":    a.structFieldName,
		"name":     a.fieldName,
		"funcName": funcName,
		"embed":    a.embed,
	}

	return returnWithTemplate("call_custom", data)
}

func getCustomArrayFunc(a ruleFuncArgs) (string, error) {
	rd.Imports["strconv"] = struct{}{}

	funcName := strings.TrimPrefix(a.rule, "@")

	if funcName == "" {
		funcName = "Validate"
	}

	data := map[string]any{
		"field":    a.structFieldName,
		"name":     a.fieldName,
		"funcName": funcName,
	}

	return returnWithTemplate("call_custom_array", data)
}
