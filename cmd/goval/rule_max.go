package main

func getFloat64Max(structFieldName, fieldName, rule string) (string, error) {
	v, err := getFloat64AfterEq(rule)
	if err != nil {
		return "", err
	}

	data := map[string]any{
		"field": structFieldName,
		"name":  fieldName,
		"max":   v,
	}

	return returnWithTemplate("numeric_max", data)
}

func getIntMax(structFieldName, fieldName, rule string) (string, error) {
	v, err := getIntAfterEq(rule)
	if err != nil {
		return "", err
	}

	data := map[string]any{
		"field": structFieldName,
		"name":  fieldName,
		"max":   v,
	}

	return returnWithTemplate("numeric_max", data)
}

func getStringMax(structFieldName, fieldName, rule string) (string, error) {
	v, err := getIntAfterEq(rule)
	if err != nil {
		return "", err
	}

	data := map[string]any{
		"field": structFieldName,
		"name":  fieldName,
		"max":   v,
	}

	return returnWithTemplate("string_max", data)
}
