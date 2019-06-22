package piscine
import "strings"
const nbr = "0123456789"

func Any(f func(string) bool, arr []string) bool {
	var boo bool
	for _, val := range arr {

		if f(val) &&  strings.Contains("4", val) {
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
func Nume(str string) {



}