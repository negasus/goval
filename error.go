package goval

import (
	"fmt"
	"strings"
)

// Errors is a map of field names to a slice of errors
type Errors map[string][]Error

// Add adds an error to the Errors map
func (e Errors) Add(fieldName string, err Error) {
	e[fieldName] = append(e[fieldName], err)
}

// NewCustomError creates a new custom error
func NewCustomError(messageID int) Error {
	return Error{
		Type:            ErrorTypeCustom,
		Values:          make(map[string]any),
		customMessageID: messageID,
	}
}

// NewError creates a new error
func NewError(t ErrorType) Error {
	return Error{
		Type:   t,
		Values: make(map[string]any),
	}
}

// AddValue adds a value to the error
func (e Error) AddValue(key string, value any) Error {
	e.Values[key] = value
	return e
}

// Error is a struct that represents an error
type Error struct {
	Type            ErrorType
	Values          map[string]any
	customMessageID int
}

func (e Error) customMessageStringLang(ln string) string {
	lang, ok := customMessages[e.customMessageID]
	if !ok {
		return "Unknown custom message"
	}

	s, ok := lang[ln]
	if !ok {
		return "Undefined lang for custom message"
	}

	for k, v := range e.Values {
		s = strings.Replace(s, "{"+k+"}", fmt.Sprintf("%v", v), -1)
	}

	return s
}

// StringLang returns the error message in the specified language
func (e Error) StringLang(ln string) string {
	if e.customMessageID > 0 {
		return e.customMessageStringLang(ln)
	}

	lang, ok := langs[ln]
	if !ok {
		lang = langEn
	}

	s, ok := lang[e.Type]
	if !ok {
		return "Unknown error"
	}

	for k, v := range e.Values {
		s = strings.Replace(s, "{"+k+"}", fmt.Sprintf("%v", v), -1)
	}

	return s
}

// String returns the error message in the default language
func (e Error) String() string {
	return e.StringLang(defaultLang)
}
