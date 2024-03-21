package main

import (
	"fmt"
	"go/ast"
)

type ruleFunc func(structFieldName, fieldName, rule string) (string, error)

type entry struct {
	Dir         string
	Filename    string
	PackageName string
	StructName  string
	Struct      *ast.StructType
	fields      []field
}

type rule struct {
	value string
	fn    ruleFunc
}

type field struct {
	structFieldName string
	fieldName       string
	rules           []rule
}

func (e *entry) renderBlocks() ([]string, error) {
	var res []string
	for _, f := range e.fields {
		for _, r := range f.rules {
			v, err := r.fn(f.structFieldName, f.fieldName, r.value)
			if err != nil {
				return nil, fmt.Errorf("error render block: %w", err)
			}
			res = append(res, v)
		}
	}

	return res, nil
}
