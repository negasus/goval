package goval

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Errors struct {
	Errors []Error
}

func (e Errors) String() string {
	return e.Error()
}

func (e Errors) Error() string {
	var els []string
	for _, er := range e.Errors {
		els = append(els, er.String())
	}
	return "[" + strings.Join(els, ",") + "]"
}

func (e Errors) StringWithMessage() string {
	var els []string
	for _, er := range e.Errors {
		els = append(els, er.StringWithMessage())
	}

	return "[" + strings.Join(els, ",") + "]"
}

func (e Errors) StringWithMessageLang(ln string) string {
	var els []string
	for _, er := range e.Errors {
		els = append(els, er.StringWithMessageLang(ln))
	}

	return "[" + strings.Join(els, ",") + "]"
}

//// NewCustomError creates a new custom error
//func NewCustomError(messageID int) Error {
//	return Error{
//		Type:            ErrorTypeCustom,
//		Values:          make(map[string]any),
//		Path:            []string{},
//		customMessageID: messageID,
//	}
//}

//// NewError creates a new error
//func NewError(t ErrorType) Error {
//	return Error{
//		Type:   t,
//		Values: make(map[string]any),
//		Path:   []string{},
//	}
//}

//// AddValue adds a value to the error
//func (e Error) AddValue(key string, value any) Error {
//	e.Values[key] = value
//	return e
//}

// Error is a struct that represents an error
type Error struct {
	Type    ErrorType      `json:"type,omitempty"`
	Field   string         `json:"field,omitempty"`
	Message string         `json:"message,omitempty"`
	Values  map[string]any `json:"values,omitempty"`
	Path    []string       `json:"path,omitempty"`

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

// StringWithMessageLang returns the error message in the specified language
func (e Error) StringWithMessageLang(ln string) string {
	if e.customMessageID > 0 {
		e.Message = e.customMessageStringLang(ln)

		res, err := json.Marshal(e)
		if err != nil {
			return `{"error":"error while marshalling error"}`
		}

		return string(res)
	}

	lang, ok := langs[ln]
	if !ok {
		lang = langEn
	}

	s, ok := lang[e.Type]
	if !ok {
		return `{"error":"unknown error"}`
	}

	for k, v := range e.Values {
		s = strings.Replace(s, "{"+k+"}", fmt.Sprintf("%v", v), -1)
	}

	e.Message = s

	res, err := json.Marshal(e)
	if err != nil {
		return `{"error":"error while marshalling error"}`
	}

	return string(res)
}

func (e Error) String() string {
	res, err := json.Marshal(e)
	if err != nil {
		return `{"error":"error while marshalling error"}`
	}

	return string(res)
}

func (e Error) StringWithMessage() string {
	return e.StringWithMessageLang(defaultLang)
}
