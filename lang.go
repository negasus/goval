package goval

import "fmt"

var defaultLang = "en"

// SetDefaultLang sets the default language for the validator
func SetDefaultLang(lang string) error {
	if _, ok := langs[lang]; !ok {
		return fmt.Errorf("unknown language %s", lang)
	}

	defaultLang = lang

	return nil
}

type ErrorType string

const (
	ErrorTypeCustom     ErrorType = "custom"
	ErrorTypeMaxNumeric ErrorType = "max_numeric"
	ErrorTypeMaxString  ErrorType = "max_string_length"
	ErrorTypeMinNumeric ErrorType = "min_numeric"
	ErrorTypeMinString  ErrorType = "min_string_length"
	ErrorTypeInvalid    ErrorType = "invalid"
	ErrorTypeMinArray   ErrorType = "min_array_length"
	ErrorTypeMaxArray   ErrorType = "max_array_length"
)

// AddLanguage adds a new language to the validator
func AddLanguage(lang string, messages map[ErrorType]string) {
	langs[lang] = messages
}

var (
	langs = map[string]map[ErrorType]string{
		"en": langEn,
		"ru": langRu,
	}

	langEn = map[ErrorType]string{
		ErrorTypeMaxNumeric: "The {field} must not be greater than {rule_value}",
		ErrorTypeMaxString:  "The {field} must not be greater than {rule_value} characters",
		ErrorTypeMinNumeric: "The {field} must be at least {rule_value}",
		ErrorTypeMinString:  "The {field} must be at least {rule_value} characters",
		ErrorTypeInvalid:    "The selected {field} is invalid",
		ErrorTypeMinArray:   "The {field} must contain at least {rule_value} items",
		ErrorTypeMaxArray:   "The {field} may not contain more than {rule_value} items",
	}

	langRu = map[ErrorType]string{
		ErrorTypeMaxNumeric: "Значение поля {field} должно быть более {rule_value}",
		ErrorTypeMaxString:  "Длина поля {field} должна быть более {rule_value}",
		ErrorTypeMinNumeric: "Значение поля {field} должно быть не менее {rule_value}",
		ErrorTypeMinString:  "Длина поля {field} должна быть не менее {rule_value}",
		ErrorTypeInvalid:    "Поле {field} имеет некорректное значение",
		ErrorTypeMinArray:   "Количество элементов в поле {field} должно быть не менее {rule_value}",
		ErrorTypeMaxArray:   "Количество элементов в поле {field} должно быть не более {rule_value}",
	}
)

var customMessages = map[int]map[string]string{}

// AddCustomMessage adds a custom message for a specific message ID and language
func AddCustomMessage(messageID int, lang string, message string) {
	if _, ok := customMessages[messageID]; !ok {
		customMessages[messageID] = map[string]string{}
	}

	customMessages[messageID][lang] = message
}

func Messages(errors *Errors, lang string) error {
	for i := 0; i < len(errors.Errors); i++ {
	}
	return nil
}
