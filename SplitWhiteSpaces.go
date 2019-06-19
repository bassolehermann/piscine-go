package piscine

import "strings"

func SplitWhiteSpaces(str string) []string {
var result []string

	result = strings.SplitAfter(str, " ")
	return result
}