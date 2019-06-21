package piscine

import "strings"

const num = "0123456789"

func Any(f func(string) bool, arr []string) bool {
	var boo bool
	for _, val := range arr {

		if f(val) && strings.Contains(num, strings.ToLower(string(val))) {
			boo = true
		} else {
			boo = false
		}

	}
	if boo {
		return true
	} else {
		return false
	}
}
func verif(str string) {

}
