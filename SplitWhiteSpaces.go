package piscine

import "regexp"


func SplitWhiteSpaces(str string) []string {
var result []string

	re:= regexp.MustCompile (`[""]`)

	result = re. Split(str, -1)
	return result
}