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
			return func(a ruleFuncArgs) (string, error) {
				return ruleCompareLen(a.structFieldName, a.fieldName, a.rule, "<", string(goval.ErrorTypeMinArray))
			}
		}

		switch t.name {
		case "int":
			return func(a ruleFuncArgs) (string, error) {
				return ruleCompareInt(a.structFieldName, a.fieldName, a.rule, "<", string(goval.ErrorTypeMinNumeric))
			}
		case "float64":
			return func(a ruleFuncArgs) (string, error) {
				return ruleCompareFloat64(a.structFieldName, a.fieldName, a.rule, "<", string(goval.ErrorTypeMinNumeric))
			}
		case "string":
			return func(a ruleFuncArgs) (string, error) {
				return ruleCompareLen(a.structFieldName, a.fieldName, a.rule, "<", string(goval.ErrorTypeMinString))
			}
		default:
			return nil
		}
	case "max":
		if t.isArray {
			return func(a ruleFuncArgs) (string, error) {
				return ruleCompareLen(a.structFieldName, a.fieldName, a.rule, ">", string(goval.ErrorTypeMaxArray))
			}
		}

		switch t.name {
		case "int":
			return func(a ruleFuncArgs) (string, error) {
				return ruleCompareInt(a.structFieldName, a.fieldName, a.rule, ">", string(goval.ErrorTypeMaxNumeric))
			}
		case "float64":
			return func(a ruleFuncArgs) (string, error) {
				return ruleCompareFloat64(a.structFieldName, a.fieldName, a.rule, ">", string(goval.ErrorTypeMaxNumeric))
			}
		case "string":
			return func(a ruleFuncArgs) (string, error) {
				return ruleCompareLen(a.structFieldName, a.fieldName, a.rule, ">", string(goval.ErrorTypeMaxString))
			}
		default:
			return nil
		}
	case "in":
		switch t.name {
		case "int":
			return func(a ruleFuncArgs) (string, error) {
				return ruleInInt(a.structFieldName, a.fieldName, a.rule)
			}
		case "string":
			return func(a ruleFuncArgs) (string, error) {
				return ruleInString(a.structFieldName, a.fieldName, a.rule)
			}
		case "[]int":
			return func(a ruleFuncArgs) (string, error) {
				return ruleInIntSlice(a.structFieldName, a.fieldName, a.rule)
			}
		case "[]string":
			return func(a ruleFuncArgs) (string, error) {
				return ruleInStringSlice(a.structFieldName, a.fieldName, a.rule)
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
			g.embedded = true
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
		if t.isArray {
			return rule{fn: getCustomArrayFunc, value: r}, nil
		}
		return rule{fn: getCustomFunc, value: r}, nil
	}
	if t.isStar {
		// isStar supported only for custom rules
		return rule{}, fmt.Errorf("invalid rule %q for type %s\n", r, t.name)
	}

	// todo: '=' may be in string value
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
