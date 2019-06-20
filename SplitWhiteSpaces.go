package piscine

import ("strings"

		"regexp")


func SplitWhiteSpaces(str string) []string {

re:= regexp.MustCompile ( `[" ""	""\n"]` )

	return re. Split (str, -1)
}