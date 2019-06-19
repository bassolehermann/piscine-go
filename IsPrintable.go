package piscine

import "strings"

const lettre = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lettre1= "abcdefghijklmnopqrstuvwxyz"

func IsPrintable(str string) bool {
	var booleen bool

	for _, char := range str {
		if !strings.Contains(lettre, string(char)) {

			booleen = true

		} else if !strings.Contains(lettre1, string(char)) {


		}else {
			booleen = false
		}
	}

	if booleen {

		return true
	} else {

		return false
	}
}