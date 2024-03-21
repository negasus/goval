package goval

import (
	"fmt"
	"strings"
)

type Errors map[string][]Error

func (e Errors) Add(fieldName string, err Error) {
	e[fieldName] = append(e[fieldName], err)
}

func NewCustomError(messageID int) Error {
	return Error{
		Type:            ErrorTypeCustom,
		Values:          make(map[string]any),
		customMessageID: messageID,
	}
}

func NewError(t ErrorType) Error {
	return Error{
		Type:   t,
		Values: make(map[string]any),
	}
}

func (e Error) AddValue(key string, value any) Error {
	e.Values[key] = value
	return e
}

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

func (e Error) StringLang(ln string) string {
	if e.customMessageID > 0 {
		return e.customMessageStringLang(ln)
	}

	lang, ok := Langs[ln]
	if !ok {
		lang = LangEn
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

func (e Error) String() string {
	return e.StringLang(defaultLang)
}
