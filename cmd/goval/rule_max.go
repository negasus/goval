package main

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

	return returnWithTemplate("string_max", data)
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
