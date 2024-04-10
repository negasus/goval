package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"os"
	"path"
	"strings"
	"text/template"
)

type renderData struct {
	Imports     map[string]struct{}
	UsedFlags   string
	PackageName string
	Funcs       []renderDataItem
}

type renderDataItem struct {
	StructName string
	Blocks     []string
}

var opts cmdFlags

//go:embed templates/*
var templatesFS embed.FS

var (
	templates *template.Template
	rd        = renderData{
		Imports: map[string]struct{}{},
	}
)

func main() {
	errFlags := parseFlags()
	if errFlags != nil {
		flagsUsage()
		os.Exit(1)
	}

	if opts.tagName == "" {
		opts.tagName = "goval"
	}

	var errParseTemplates error
	templates, errParseTemplates = template.ParseFS(templatesFS, "templates/*.gotmpl")
	if errParseTemplates != nil {
		fmt.Printf("error parse templates, %v", errParseTemplates)
		os.Exit(1)
	}

	workDir, errGetWD := os.Getwd()
	if errGetWD != nil {
		fmt.Printf("error get work dir, %v", errGetWD)
		os.Exit(1)
	}

	entries, errScan := scanDir(workDir, opts.types)
	if errScan != nil {
		fmt.Printf("error scan workdir, %v\n", errScan)
		os.Exit(1)
	}

	for _, t := range opts.types {
		var found bool
		for _, e := range entries {
			if e.StructName == t {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("struct %q not found\n", t)
			os.Exit(1)
		}
	}

	rd.UsedFlags = strings.Join(os.Args[1:], " ")

	for _, e := range entries {
		if rd.PackageName == "" {
			rd.PackageName = e.PackageName
		}

		if rd.PackageName != e.PackageName {
			fmt.Printf("package name mismatch: %s != %s\n", rd.PackageName, e.PackageName)
			os.Exit(1)
		}

		errFillRules := fillRulesForEntry(e)
		if errFillRules != nil {
			fmt.Printf("fillRulesForEntry error: %s\n", errFillRules)
			os.Exit(1)
		}

		rdi := renderDataItem{
			StructName: e.StructName,
		}

		blocks, errBlocks := e.renderBlocks()
		if errBlocks != nil {
			fmt.Printf("renderBlocks error: %s\n", errBlocks)
			os.Exit(1)
		}

		rdi.Blocks = blocks

		rd.Funcs = append(rd.Funcs, rdi)
	}

	outFilename := fmt.Sprintf("goval_%s.go", strings.Trim(toSneakCase(strings.Join(opts.types, "_")), "_"))
	if opts.outputFilename != "" {
		outFilename = opts.outputFilename
	}

	w := bytes.NewBuffer(nil)

	errExec := templates.ExecuteTemplate(w, "main.gotmpl", rd)
	if errExec != nil {
		fmt.Printf("error exec main.gotmpl: %s\n", errExec)
		os.Exit(1)
	}

	res := bytes.TrimRight(w.Bytes(), "\n")
	res = append(res, 0x0A)

	//for i, v := range strings.Split(string(res), "\n") {
	//	fmt.Printf("%3d %s\n", i, v)
	//}

	var errFormat error
	res, errFormat = format.Source(res)
	if errFormat != nil {
		fmt.Printf("format error: %s\n", errFormat)
		os.Exit(1)
	}

	errWrite := os.WriteFile(path.Join(workDir, outFilename), res, 0644)
	if errWrite != nil {
		fmt.Printf("write error: %s\n", errWrite)
		os.Exit(1)
	}

	if opts.debug {
		fmt.Printf("write file: %s\n", path.Join(workDir, outFilename))
	}
}
