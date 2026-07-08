package iteration

import "strings"

func RepeatConcatenation(character string, iterations int) string {
	var repeated string
	for i := 0; i < iterations; i++ {
		repeated += character
	}
	return repeated
}

func RepeatBuilder(character string, iterations int) string {
	var repeated strings.Builder
	for i := 0; i < iterations; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
