package main

import (
	"strconv"
)

func ruleCompareLen(structFieldName, fieldName, rule, op, errType string) (string, error) {
	v, err := strconv.Atoi(rule)
	if err != nil {
		return "", err
	}

	return ruleCompare(structFieldName, fieldName, op, v, errType, "compare_len")
}

func ruleCompareInt(structFieldName, fieldName, rule, op, errType string) (string, error) {
	v, err := strconv.Atoi(rule)
	if err != nil {
		return "", err
	}

	return ruleCompare(structFieldName, fieldName, op, v, errType, "compare")
}

func ruleCompareFloat64(structFieldName, fieldName, rule, op, errType string) (string, error) {
	v, err := strconv.ParseFloat(rule, 64)
	if err != nil {
		return "", err
	}

	return ruleCompare(structFieldName, fieldName, op, v, errType, "compare")
}

func ruleCompare(structFieldName, fieldName, op string, v any, errType, template string) (string, error) {
	data := map[string]any{
		"field":     structFieldName,
		"name":      fieldName,
		"op":        op,
		"value":     v,
		"errorType": errType,
		"meta": map[string]any{
			"rule_value": v,
		},
	}

	return returnWithTemplate(template, data)
}
