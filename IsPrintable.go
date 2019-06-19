package piscine

import "strings"

const lettre = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"


func IsPrintable(str string) bool {
	var booleen bool

	for _, char := range str {
		if strings.Contains(lettre, string(char)) {

			booleen = true

		} else {
			booleen = false
		}
	}

	if booleen {

		return true
	} else {

		return false
	}
}
