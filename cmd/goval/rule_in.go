package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/negasus/goval"
)

func ruleInSlice(structFieldName, fieldName, sliceType, v, templateName string) (string, error) {
	rd.Imports["slices"] = struct{}{}

	data := map[string]any{
		"field":     structFieldName,
		"name":      fieldName,
		"sliceType": sliceType,
		"var":       v,
		"errorType": goval.ErrorTypeInvalid,
		"meta":      map[string]any{},
	}

	return returnWithTemplate(templateName, data)
}

func isRuleVar(rule string) (string, bool) {
	if len(rule) < 3 {
		return "", false
	}

	if rule[0] == '{' && rule[len(rule)-1] == '}' {
		return rule[1 : len(rule)-1], true
	}

	return "", false
}

func ruleInStringSlice(structFieldName, fieldName, rule string) (string, error) {
	// in={varName}
	if v, ok := isRuleVar(rule); ok {
		return ruleInSlice(structFieldName, fieldName, "", v, "in_slice")
	}

	// in=a,c,b
	return ruleInSlice(structFieldName, fieldName, "[]string", makeStringFromStrings(parseStringInput(rule)), "in_slice")
}

func ruleInIntSlice(structFieldName, fieldName, rule string) (string, error) {
	// in={varName}
	if v, ok := isRuleVar(rule); ok {
		return ruleInSlice(structFieldName, fieldName, "", v, "in_slice")
	}

	var ii []int

	for _, i := range strings.Split(rule, ",") {
		v, ok := strconv.Atoi(i)
		if ok != nil {
			return "", fmt.Errorf("invalid supported syntax for rule 'in', %s", rule)
		}
		ii = append(ii, v)
	}

	// in=1,2,3
	return ruleInSlice(structFieldName, fieldName, "[]int", makeStringFromInts(ii), "in_slice")
}

func ruleInString(structFieldName, fieldName, rule string) (string, error) {
	// in={varName}
	if v, ok := isRuleVar(rule); ok {
		return ruleInSlice(structFieldName, fieldName, "", v, "in")
	}

	// in=a,c,b
	return ruleInSlice(structFieldName, fieldName, "[]string", makeStringFromStrings(parseStringInput(rule)), "in")
}

func ruleInInt(structFieldName, fieldName, rule string) (string, error) {
	// in={varName}
	if v, ok := isRuleVar(rule); ok {
		return ruleInSlice(structFieldName, fieldName, "", v, "in")
	}

	var ii []int

	for _, i := range strings.Split(rule, ",") {
		v, ok := strconv.Atoi(i)
		if ok != nil {
			return "", fmt.Errorf("invalid supported syntax for rule 'in', %s", rule)
		}
		ii = append(ii, v)
	}

	// in=1,2,3
	return ruleInSlice(structFieldName, fieldName, "[]int", makeStringFromInts(ii), "in")
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
