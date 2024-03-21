package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"strings"
)

func scanDir(workDir string, srcType []string) ([]*entry, error) {
	dirData, errReadDir := os.ReadDir(workDir)
	if errReadDir != nil {
		return nil, errReadDir
	}

	var entries []*entry

	for _, e := range dirData {
		if e.IsDir() {
			continue
		}

		if !strings.HasSuffix(e.Name(), ".go") {
			continue
		}

		ee, errParse := parseFile(workDir, e.Name(), srcType)
		if errParse != nil {
			return nil, errParse
		}

		entries = append(entries, ee...)
	}

	return entries, nil
}

func parseFile(dir, filename string, srcType []string) ([]*entry, error) {
	if opts.Debug {
		fmt.Printf("...parse file: %s\n", path.Join(dir, filename))
	}

	parsedData, errParse := parser.ParseFile(token.NewFileSet(), path.Join(dir, filename), nil, 0)
	if errParse != nil {
		return nil, errParse
	}

	var entries []*entry

	for _, d := range parsedData.Decls {
		switch v := d.(type) {
		case *ast.GenDecl:
			if v.Tok != token.TYPE {
				continue
			}
			if len(v.Specs) == 0 {
				continue
			}

			s1, ok := v.Specs[0].(*ast.TypeSpec)
			if !ok {
				continue
			}

			typ, okT := s1.Type.(*ast.StructType)
			if !okT {
				continue
			}

			if typ.Fields == nil {
				continue
			}

			if len(typ.Fields.List) == 0 {
				continue
			}

			if s1.Name == nil {
				continue
			}

			for _, t := range srcType {
				if t == s1.Name.Name {
					e := &entry{
						Dir:        dir,
						Filename:   filename,
						StructName: s1.Name.Name,
						Struct:     typ,
					}
					if parsedData.Name != nil {
						e.PackageName = parsedData.Name.Name
					}

					entries = append(entries, e)
				}
			}

		}
	}

	return entries, nil
}
