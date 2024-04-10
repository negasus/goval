package main

import (
	"fmt"
	"go/ast"
)

type ruleFuncArgs struct {
	structFieldName string
	fieldName       string
	rule            string
	embed           bool
}

type ruleFunc func(args ruleFuncArgs) (string, error)

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
	embedded        bool
}

func (e *entry) renderBlocks() ([]string, error) {
	var res []string
	for _, f := range e.fields {
		for _, r := range f.rules {
			a := ruleFuncArgs{
				structFieldName: f.structFieldName,
				fieldName:       f.fieldName,
				rule:            r.value,
				embed:           f.embedded,
			}
			v, err := r.fn(a)
			if err != nil {
				return nil, fmt.Errorf("error render block: %w", err)
			}
			res = append(res, v)
		}
	}

	return res, nil
}
