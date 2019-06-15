package piscine

import "unicode/utf8"

func StrLen(str string) int {

	return utf8.RuneCountInString(str)

	 

}