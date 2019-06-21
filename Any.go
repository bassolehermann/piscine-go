package piscine

func Any(f func(string) bool, arr []string) bool {
	var boo bool
	for _, val := range arr {

		if f(val) {
			boo = false
		} else {
			boo = true
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
