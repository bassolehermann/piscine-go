package piscine



func CountIf(f func(string) bool, tab []string) int {
	var boo bool
	var nombre int
	for i, val := range tab {

		if f(val) {
			boo = true
			nombre=i+1
		} else {
			boo = false
			nombre =(i-len(val))
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