package piscine

import "regexp"


func SplitWhiteSpaces(str string) []string {

re:= regexp.MustCompile ( `[" ""	"]` )

	return re. Split (str, -1)
}