package piscine

import "strings"


func SplitWhiteSpaces(str string) []string {
var result []string

	result = strings.Split(str," 	\n")
	return result
}