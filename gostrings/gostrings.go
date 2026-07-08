package gostrings

import "strings"

func Compare(a, b string) int {
	return strings.Compare(a, b)
}

func Contains(a, b string) bool {
	return strings.Contains(a, b)
}
