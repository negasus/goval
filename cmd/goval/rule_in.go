package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getIntIn(structFieldName, fieldName, rule string) (string, error) {
	v := rule[3:]
	if len(v) == 0 {
		return "", fmt.Errorf("invalid supported syntax for rule 'in', %s", rule)
	}

	// in={varName}
	if v[0] == '{' {
		if len(v) == 1 {
			return "", fmt.Errorf("invalid supported syntax for rule 'in', %s", rule)
		}
		if v[len(v)-1] != '}' {
			return "", fmt.Errorf("invalid supported syntax for rule 'in', %s", rule)
		}
		data := map[string]any{
			"field": structFieldName,
			"name":  fieldName,
			"var":   v[1 : len(v)-1],
		}

		return returnWithTemplate("string_in_var", data)
	}

	var ii []int

	for _, i := range strings.Split(v, ",") {
		v, ok := strconv.Atoi(i)
		if ok != nil {
			return "", fmt.Errorf("invalid supported syntax for rule 'in', %s", rule)
		}
		ii = append(ii, v)
	}

	// in=1,2,3
	data := map[string]any{
		"field": structFieldName,
		"name":  fieldName,
		"var":   makeStringFromInts(ii),
	}
	return returnWithTemplate("int_enum", data)
}

func getStringIn(structFieldName, fieldName, rule string) (string, error) {
	v := rule[3:]
	if len(v) == 0 {
		return "", fmt.Errorf("invalid supported syntax for rule 'in', %s", rule)
	}

	// in={varName}
	if v[0] == '{' {
		if len(v) == 1 {
			return "", fmt.Errorf("invalid supported syntax for rule 'in', %s", rule)
		}
		if v[len(v)-1] != '}' {
			return "", fmt.Errorf("invalid supported syntax for rule 'in', %s", rule)
		}
		data := map[string]any{
			"field": structFieldName,
			"name":  fieldName,
			"var":   v[1 : len(v)-1],
		}

		return returnWithTemplate("string_in_var", data)
	}

	// in=foo,bar
	data := map[string]any{
		"field": structFieldName,
		"name":  fieldName,
		"var":   makeStringFromStrings(parseStringInput(v)),
	}
	return returnWithTemplate("string_enum", data)
}

func parseStringInput(input string) []string {
	var result []string
	var current strings.Builder
	var inQuotes, isEscaped bool

	for _, r := range input {
		if isEscaped {
			current.WriteRune(r)
			isEscaped = false
			continue
		}

		switch r {
		case '\\':
			isEscaped = true
		case '\'':
			inQuotes = !inQuotes
		case ',':
			if inQuotes {
				current.WriteRune(r)
			} else {
				result = append(result, current.String())
				current.Reset()
			}
		default:
			current.WriteRune(r)
		}
	}

	result = append(result, current.String())

	return result
}

func makeStringFromInts(slice []int) string {
	var vv []string
	for _, v := range slice {
		vv = append(vv, fmt.Sprintf("%d", v))
	}
	return "{" + strings.Join(vv, ", ") + "}"
}

func makeStringFromStrings(slice []string) string {
	for i, v := range slice {
		slice[i] = fmt.Sprintf("%q", v)
	}
	return "{" + strings.Join(slice, ", ") + "}"
}
