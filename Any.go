package piscine


import "strings"


func Any(f func(string) bool, arr []string) bool {
	var boo bool
	for _, val := range arr {

		if f(val) {
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
func Nume(str string) bool{

 	if !strings.Contains(str, "") {

	return true

	} else{
	return false
	
	}

}