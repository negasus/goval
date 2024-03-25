package main

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/negasus/goval"
)

func getRuleFunc(funcName, typeName string) ruleFunc {
	switch funcName {
	case "min":
		switch typeName {
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
		case "[]string", "[]int":
			return func(structFieldName, fieldName, rule string) (string, error) {
				return ruleCompareLen(structFieldName, fieldName, rule, "<", goval.ErrorTypeMinArray.String())
			}
		default:
			return nil
		}
	case "max":
		switch typeName {
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
		case "[]string", "[]int":
			return func(structFieldName, fieldName, rule string) (string, error) {
				return ruleCompareLen(structFieldName, fieldName, rule, ">", goval.ErrorTypeMaxArray.String())
			}
		default:
			return nil
		}
	case "in":
		switch typeName {
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

func getTypeName(f ast.Expr) (string, bool) {
	switch v := f.(type) {
	case *ast.ArrayType:
		s, ok := getTypeName(v.Elt)
		if !ok {
			return "", false
		}
		return "[]" + s, true
	case *ast.Ident:
		return v.Name, true
	default:
		return "", false
	}
}

func fillRulesForEntry(e *entry) error {
	for _, f := range e.Struct.Fields.List {
		if f.Tag == nil {
			continue
		}

		typeName, okTypeName := getTypeName(f.Type)
		if !okTypeName {
			continue
		}

		rules := getGovalRules(f.Tag.Value)
		if len(rules) == 0 {
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
			ri, err := parseRule(r, typeName)
			if err != nil {
				return fmt.Errorf("error parse rule: %w", err)
			}

			g.rules = append(g.rules, ri)

		}

		e.fields = append(e.fields, g)
	}

	return nil
}

func parseRule(r string, typeName string) (rule, error) {
	if strings.HasPrefix(r, "@") {
		return rule{fn: getCustomFunc, value: r}, nil
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
		fn:    getRuleFunc(parts[0], typeName),
	}

	if fr.fn == nil {
		return rule{}, fmt.Errorf("invalid rule %q for type %s\n", r, typeName)
	}

	return fr, nil
}
