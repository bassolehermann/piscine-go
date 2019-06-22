package piscine

import "strings"

const nbr = "0123456789"

func Any(f func(string) bool, arr []string) bool {
	var boo bool
	for i, val := range arr {

		if f(val) && !strings.Contains(nbr, strings.ToLower(string(val[i]))) {
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
