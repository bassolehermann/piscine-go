
const nbr = "0123456789"

func Any(f func(string) bool, arr []string) bool {
	var boo bool
	for _, val := range arr {

		if f(val) {
			boo = true
		} else if f(val){
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
