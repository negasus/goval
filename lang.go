package goval

import "fmt"

var defaultLang = "en"

func SetDefaultLang(lang string) error {
	if _, ok := Langs[lang]; !ok {
		return fmt.Errorf("unknown language %s", lang)
	}

	defaultLang = lang
	return nil
}

type ErrorType int

const (
	ErrorTypeCustom ErrorType = iota
	ErrorTypeIn
	ErrorTypeMaxNumeric
	ErrorTypeMaxString
	ErrorTypeMinNumeric
	ErrorTypeMinString
	ErrorTypeInvalid
)

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
	default:
		return "Unknown error type"
	}
}

var (
	Langs = map[string]map[ErrorType]string{
		"en": LangEn,
		"ru": LangRu,
	}

	LangEn = map[ErrorType]string{
		ErrorTypeIn:         "The selected {field} is invalid",
		ErrorTypeMaxNumeric: "The {field} must not be greater than {max}",
		ErrorTypeMaxString:  "The {field} must not be greater than {max} characters",
		ErrorTypeMinNumeric: "The {field} must be at least {min}",
		ErrorTypeMinString:  "The {field} must be at least {min} characters",
		ErrorTypeInvalid:    "The selected {field} is invalid",
	}

	LangRu = map[ErrorType]string{
		ErrorTypeIn:         "Поле {field} имеет некорректное значение",
		ErrorTypeMaxNumeric: "Поле {field} должно быть больше {max}",
		ErrorTypeMaxString:  "Длина поля {field} должна быть более {max}",
		ErrorTypeMinNumeric: "Поле {field} должно быть не менее {min}",
		ErrorTypeMinString:  "Длина поля {field} должна быть не менее {min}",
		ErrorTypeInvalid:    "Поле {field} имеет некорректное значение",
	}
)

var customMessages = map[int]map[string]string{}

func AddCustomMessage(messageID int, lang string, message string) {
	if _, ok := customMessages[messageID]; !ok {
		customMessages[messageID] = map[string]string{}
	}

	customMessages[messageID][lang] = message
}
