package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var reGovalTags *regexp.Regexp

func getGovalRules(s string) []string {
	if reGovalTags == nil {
		reGovalTags = regexp.MustCompile(opts.tagName + `:"([^"]+)"`)
	}

	matches := reGovalTags.FindStringSubmatch(s)
	if len(matches) > 1 {
		return strings.Split(matches[1], ";")
	}

	return nil
}

var reJsonTags = regexp.MustCompile(`json:"([^"]+)"`)

func getJsonName(s string) string {
	matches := reJsonTags.FindStringSubmatch(s)
	if len(matches) > 1 {
		return strings.TrimSuffix(matches[1], ",omitempty")
	}

	return ""
}

func toSneakCase(s string) string {
	var res string
	for i, r := range s {
		if i > 0 && 'A' <= r && r <= 'Z' {
			res += "_"
		}
		res += string(r)
	}
	return strings.ToLower(res)
}

func returnWithTemplate(tmplName string, data map[string]any) (string, error) {
	buf := bytes.NewBuffer(nil)

	errExec := templates.ExecuteTemplate(buf, tmplName+".gotmpl", data)
	if errExec != nil {
		return "", fmt.Errorf("error execute template, %w", errExec)
	}

	return buf.String(), nil
}
