package main

func getFloat64Min(structFieldName, fieldName, rule string) (string, error) {
	v, err := getFloat64AfterEq(rule)
	if err != nil {
		return "", err
	}

	data := map[string]any{
		"field": structFieldName,
		"name":  fieldName,
		"min":   v,
	}

	return returnWithTemplate("numeric_min", data)
}

func getIntMin(structFieldName, fieldName, rule string) (string, error) {
	v, err := getIntAfterEq(rule)
	if err != nil {
		return "", err
	}

	data := map[string]any{
		"field": structFieldName,
		"name":  fieldName,
		"min":   v,
	}

	return returnWithTemplate("numeric_min", data)
}

func getStringMin(structFieldName, fieldName, rule string) (string, error) {
	v, err := getIntAfterEq(rule)
	if err != nil {
		return "", err
	}

	data := map[string]any{
		"field": structFieldName,
		"name":  fieldName,
		"min":   v,
	}

	return returnWithTemplate("string_min", data)
}
