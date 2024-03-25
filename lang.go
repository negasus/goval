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

// ErrorType is an enum for the different types of errors that can be returned by the validator
type ErrorType int

const (
	ErrorTypeCustom ErrorType = iota
	ErrorTypeIn
	ErrorTypeMaxNumeric
	ErrorTypeMaxString
	ErrorTypeMinNumeric
	ErrorTypeMinString
	ErrorTypeInvalid
	ErrorTypeMinArray
	ErrorTypeMaxArray
)

// String returns the string representation of the error type
func (t ErrorType) String() string {
	switch t {
	case ErrorTypeIn:
		return "ErrorTypeIn"
	case ErrorTypeMaxNumeric:
		return "ErrorTypeMaxNumeric"
	case ErrorTypeMaxString:
		return "ErrorTypeMaxString"
	case ErrorTypeMinNumeric:
		return "ErrorTypeMinNumeric"
	case ErrorTypeMinString:
		return "ErrorTypeMinString"
	case ErrorTypeInvalid:
		return "ErrorTypeInvalid"
	case ErrorTypeMinArray:
		return "ErrorTypeMinArray"
	case ErrorTypeMaxArray:
		return "ErrorTypeMaxArray"
	default:
		return fmt.Sprintf("Unknown error type: '%d'", t)
	}
}

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
		ErrorTypeIn:         "The selected {field} is invalid",
		ErrorTypeMaxNumeric: "The {field} must not be greater than {rule_value}",
		ErrorTypeMaxString:  "The {field} must not be greater than {rule_value} characters",
		ErrorTypeMinNumeric: "The {field} must be at least {rule_value}",
		ErrorTypeMinString:  "The {field} must be at least {rule_value} characters",
		ErrorTypeInvalid:    "The selected {field} is invalid",
		ErrorTypeMinArray:   "The {field} must contain at least {rule_value} items",
		ErrorTypeMaxArray:   "The {field} may not contain more than {rule_value} items",
	}

	langRu = map[ErrorType]string{
		ErrorTypeIn:         "Поле {field} имеет некорректное значение",
		ErrorTypeMaxNumeric: "Поле {field} должно быть больше {rule_value}",
		ErrorTypeMaxString:  "Длина поля {field} должна быть более {rule_value}",
		ErrorTypeMinNumeric: "Поле {field} должно быть не менее {rule_value}",
		ErrorTypeMinString:  "Длина поля {field} должна быть не менее {rule_value}",
		ErrorTypeInvalid:    "Поле {field} имеет некорректное значение",
		ErrorTypeMinArray:   "Поле {field} должно содержать не менее {rule_value} элементов",
		ErrorTypeMaxArray:   "Поле {field} должно содержать не более {rule_value} элементов",
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
