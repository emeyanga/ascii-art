package main

import (
	"strings"
)

func Generate(input string, banner map[rune][]string) string {
	if input == "\\n" {
		return ""
	}

	segments := SplitInput(input)
	var output []string

	for _, seg := range segments {
		if seg == "" {
			output = append(output, "")
		} else {
			rows := RenderLine(seg, banner)
			output = append(output, rows...)
		}
	}
	return strings.Join(output, "\n")
}
