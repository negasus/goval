package main

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/negasus/goval"
)

func getRuleFunc(funcName string, t *typeName) ruleFunc {
	switch funcName {
	case "min":
		if t.isArray {
			return func(structFieldName, fieldName, rule string) (string, error) {
				return ruleCompareLen(structFieldName, fieldName, rule, "<", goval.ErrorTypeMinArray.String())
			}
		}

		switch t.name {
		case "int":
			return func(structFieldName, fieldName, rule string) (string, error) {
				return ruleCompareInt(structFieldName, fieldName, rule, "<", goval.ErrorTypeMinNumeric.String())
			}
		case "float64":
			return func(structFieldName, fieldName, rule string) (string, error) {
				return ruleCompareFloat64(structFieldName, fieldName, rule, "<", goval.ErrorTypeMinNumeric.String())
			}
		case "string":
			return func(structFieldName, fieldName, rule string) (string, error) {
				return ruleCompareLen(structFieldName, fieldName, rule, "<", goval.ErrorTypeMinString.String())
			}
		default:
			return nil
		}
	case "max":
		if t.isArray {
			return func(structFieldName, fieldName, rule string) (string, error) {
				return ruleCompareLen(structFieldName, fieldName, rule, ">", goval.ErrorTypeMaxArray.String())
			}
		}

		switch t.name {
		case "int":
			return func(structFieldName, fieldName, rule string) (string, error) {
				return ruleCompareInt(structFieldName, fieldName, rule, ">", goval.ErrorTypeMaxNumeric.String())
			}
		case "float64":
			return func(structFieldName, fieldName, rule string) (string, error) {
				return ruleCompareFloat64(structFieldName, fieldName, rule, ">", goval.ErrorTypeMaxNumeric.String())
			}
		case "string":
			return func(structFieldName, fieldName, rule string) (string, error) {
				return ruleCompareLen(structFieldName, fieldName, rule, ">", goval.ErrorTypeMaxString.String())
			}
		default:
			return nil
		}
	case "in":
		switch t.name {
		case "int":
			return func(structFieldName, fieldName, rule string) (string, error) {
				return ruleInInt(structFieldName, fieldName, rule)
			}
		case "string":
			return func(structFieldName, fieldName, rule string) (string, error) {
				return ruleInString(structFieldName, fieldName, rule)
			}
		case "[]int":
			return func(structFieldName, fieldName, rule string) (string, error) {
				return ruleInIntSlice(structFieldName, fieldName, rule)
			}
		case "[]string":
			return func(structFieldName, fieldName, rule string) (string, error) {
				return ruleInStringSlice(structFieldName, fieldName, rule)
			}
		default:
			return nil
		}
	default:
		return nil
	}
}

func fillRulesForEntry(e *entry) error {
	for _, f := range e.Struct.Fields.List {
		if f.Tag == nil {
			continue
		}

		rules := getGovalRules(f.Tag.Value)
		if len(rules) == 0 {
			continue
		}

		tn := getTypeName(f.Type)
		if tn == nil {
			continue
		}

		g := field{}

		if len(f.Names) > 0 {
			// normal field
			g.structFieldName = f.Names[0].Name
		} else {
			// embedded struct
			i, okI := f.Type.(*ast.Ident)
			if !okI {
				continue
			}
			g.structFieldName = i.Name
		}

		g.fieldName = getJsonName(f.Tag.Value)

		if g.fieldName == "" {
			g.fieldName = g.structFieldName
		}

		for _, r := range rules {
			ri, err := parseRule(r, tn)
			if err != nil {
				return fmt.Errorf("error parse rule: %w", err)
			}

			g.rules = append(g.rules, ri)
		}

		e.fields = append(e.fields, g)
	}

	return nil
}

func parseRule(r string, t *typeName) (rule, error) {
	if strings.HasPrefix(r, "@") {
		return rule{fn: getCustomFunc, value: r}, nil
	}
	if t.isStar {
		// isStar supported only for custom rules
		return rule{}, fmt.Errorf("invalid rule %q for type %s\n", r, t.name)
	}

	parts := strings.Split(r, "=")
	if len(parts) > 2 {
		return rule{}, fmt.Errorf("invalid rule %q\n", r)
	}
	if parts[0] == "" {
		return rule{}, fmt.Errorf("invalid rule %q\n", r)
	}

	var tail string
	if len(parts) == 2 {
		tail = parts[1]
	}

	fr := rule{
		value: tail,
		fn:    getRuleFunc(parts[0], t),
	}

	if fr.fn == nil {
		return rule{}, fmt.Errorf("invalid rule %q for type %q\n", r, t.name)
	}

	return fr, nil
}
