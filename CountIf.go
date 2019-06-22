package piscine



func CountIf(f func(string) bool, tab []string) int {
	var boo bool
	var nombre int
	for i, val := range tab {

		if f(val) {
			boo = true
			nombre=len(val)-i
		} else {
			boo = false
			nombre =len(val)
		}

	}
	if boo {
		return nombre
	} else {
		return nombre

	}
}
func count(str string) {



}