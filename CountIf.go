package piscine



func CountIf(f func(string) bool, tab []string) int {
	var boo bool
	var nombre int
	for i, val := range tab {

		if f(val) {
			boo = true
			nombre=i
		} else if f(val){
			boo = false
			nombre =i
		}

	}
	if boo {
		return nombre
	} else {
		return nombre

	}
}
func Nume(str string) {



}