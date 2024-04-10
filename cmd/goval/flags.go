package main

import (
	"fmt"
	"os"
)

type cmdFlags struct {
	debug          bool
	types          []string
	outputFilename string
	tagName        string
}

func flagsUsage() {
	fmt.Println("Usage: goval -t <type> [-t <type>] [-d] [-o <output filename>] [-n <tag name>]")
}

func parseFlags() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("no flags provided")
	}
	var i int
	for {
		i++
		if i >= len(os.Args) {
			break
		}
		switch os.Args[i] {
		case "-d", "--debug":
			if opts.debug {
				return fmt.Errorf("debug flag already provided")
			}
			opts.debug = true
			continue
		case "-t", "--type":
			i++
			if i >= len(os.Args) {
				return fmt.Errorf("no types provided")
			}
			opts.types = append(opts.types, os.Args[i])
			continue
		case "-o", "--output":
			if opts.outputFilename != "" {
				return fmt.Errorf("output filename already provided")
			}
			i++
			if i >= len(os.Args) {
				return fmt.Errorf("no output filename provided")
			}
			opts.outputFilename = os.Args[i]
			continue
		case "-n", "--tag":
			if opts.tagName != "" {
				return fmt.Errorf("tag name already provided")
			}
			i++
			if i >= len(os.Args) {
				return fmt.Errorf("no tag name provided")
			}
			opts.tagName = os.Args[i]
			continue
		default:
			return fmt.Errorf("unknown flag %q", os.Args[i])
		}
	}

	if len(opts.types) == 0 {
		return fmt.Errorf("no types provided")
	}

	return nil
}
