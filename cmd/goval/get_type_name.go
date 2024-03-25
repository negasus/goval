package main

import "go/ast"

type typeName struct {
	name    string
	isArray bool
	isStar  bool
}

func getTypeName(f ast.Expr, depth ...int) *typeName {
	t := typeName{}
	switch v := f.(type) {
	case *ast.ArrayType:
		t.isArray = true

		if len(depth) == 0 {
			t2 := getTypeName(v.Elt, append(depth, 1)...)
			if t2 != nil {
				t.name = "[]" + t2.name
			}
		}

		return &t
	case *ast.Ident:
		t.name = v.Name
		return &t
	case *ast.StarExpr:
		t.isStar = true
		t2 := getTypeName(v.X)
		if t2 != nil {
			t.name = t2.name
		}
		return &t
	default:
		return nil
	}
}
