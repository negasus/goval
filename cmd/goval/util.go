package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	reTags *regexp.Regexp
)

func getGovalRules(s string) []string {
	if reTags == nil {
		reTags = regexp.MustCompile(opts.TagName + `:"([^"]+)"`)
	}

	matches := reTags.FindStringSubmatch(s)
	if len(matches) > 1 {
		return strings.Split(matches[1], ";")
	}

	return nil
}

var jsonTags = regexp.MustCompile(`json:"([^"]+)"`)

func getJsonName(s string) string {
	matches := jsonTags.FindStringSubmatch(s)
	if len(matches) > 1 {
		return strings.TrimSuffix(matches[1], ",omitempty")
	}

	return ""
}

func getIntAfterEq(rule string) (int, error) {
	parts := strings.Split(rule, "=")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid rule: %s", rule)
	}

	v, errConvert := strconv.ParseInt(parts[1], 10, 64)
	if errConvert != nil {
		return 0, fmt.Errorf("convert min: %w", errConvert)
	}

	return int(v), nil
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
