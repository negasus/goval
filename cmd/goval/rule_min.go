package main

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

	return returnWithTemplate("int_min", data)
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
