package piscine

const nbr = "0123456789"

func CountIf(f func(string) bool, tab []string) int {
	var boo bool
	var nbr int
	for i, val := range tab {

		if f(val) {
			boo = true
			nbr=i
		} else if f(val){
			boo = false
			nbr =i
		}

	}
	if boo {
		return nbr
	} else {
		return nbr

	}
}
func Nume(str string) {



}