package main

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/umputun/go-flags"
)

type RenderData struct {
	UsedFlags   string
	PackageName string
	Funcs       []RenderDataItem
}

type RenderDataItem struct {
	StructName string
	Blocks     []string
}

//go:embed templates/*
var templatesFS embed.FS

var (
	templates *template.Template
)

var opts struct {
	Debug          bool     `short:"d" long:"debug" description:"Debug mode"`
	SrcType        []string `short:"t" long:"type" description:"Type, multiple allowed" required:"true"`
	OutputFilename string   `short:"o" long:"output" description:"Output file"`
	TagName        string   `short:"n" long:"tag" description:"Tag name" default:"goval"`
}

func main() {
	p := flags.NewParser(&opts, flags.PrintErrors|flags.PassDoubleDash|flags.HelpFlag)
	_, err := p.Parse()
	if err != nil {
		os.Exit(2)
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

	entries, errScan := scanDir(workDir, opts.SrcType)
	if errScan != nil {
		fmt.Printf("error scan workdir, %v\n", errScan)
		os.Exit(1)
	}

	for _, t := range opts.SrcType {
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

	rd := RenderData{
		UsedFlags: strings.Join(os.Args[1:], " "),
	}

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

		rdi := RenderDataItem{
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

	outFilename := fmt.Sprintf("goval_%s.go", strings.Trim(toSneakCase(strings.Join(opts.SrcType, "_")), "_"))
	if opts.OutputFilename != "" {
		outFilename = opts.OutputFilename
	}

	w := bytes.NewBuffer(nil)

	errExec := templates.ExecuteTemplate(w, "main.gotmpl", rd)
	if errExec != nil {
		fmt.Printf("error exec main.gotmpl: %s\n", errExec)
		os.Exit(1)
	}

	res := bytes.TrimRight(w.Bytes(), "\n")
	res = append(res, 0x0A)

	errWrite := os.WriteFile(path.Join(workDir, outFilename), res, 0644)
	if errWrite != nil {
		fmt.Printf("write error: %s\n", errWrite)
		os.Exit(1)
	}

	if opts.Debug {
		fmt.Printf("write file: %s\n", path.Join(workDir, outFilename))
	}
}
