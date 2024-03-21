package main

import (
	"fmt"
	"go/ast"
	"strings"
)

var ruleFuncs = map[string]map[string]ruleFunc{
	"@": {
		"@": getCustomFunc,
	},
	"min": {
		"float64": getFloat64Min,
		"int":     getIntMin,
		"string":  getStringMin,
	},
	"max": {
		"float64": getFloat64Max,
		"int":     getIntMax,
		"string":  getStringMax,
	},
	"in": {
		"int":    getIntIn,
		"string": getStringIn,
	},
}

func fillRulesForEntry(e *entry) error {
	for _, f := range e.Struct.Fields.List {
		if f.Tag == nil {
			continue
		}

		typ, okTp := f.Type.(*ast.Ident)
		if !okTp {
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
			ri, err := parseRule(r, typ.Name)
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
	fr := rule{
		value: r,
	}

	var ok bool

	if strings.HasPrefix(r, "@") {
		fr.fn = getCustomFunc
		return fr, nil
	}
	if strings.HasPrefix(r, "min=") {
		fr.fn, ok = ruleFuncs["min"][typeName]
	}
	if strings.HasPrefix(r, "max=") {
		fr.fn, ok = ruleFuncs["max"][typeName]
	}
	if strings.HasPrefix(r, "in=") {
		fr.fn, ok = ruleFuncs["in"][typeName]
	}

	if !ok {
		return fr, fmt.Errorf("unsupported type %q for rule %q\n", typeName, r)
	}

	return fr, nil
}
