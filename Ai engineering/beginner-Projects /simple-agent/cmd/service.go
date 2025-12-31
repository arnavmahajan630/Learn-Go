package cmd

import "strings"

func CleanInput(str string) []string {
	lowered := strings.ToLower(str)  // lowersize all text
	words := strings.Fields(lowered) // break down string into words for messy input
	return words
}