package piscine

func Rot14(str string) string {
	var r rune
	r = []rune(str)

	if r >= 'a' && r <= 'z' {

		if r >= 'm' {
			return r - 14
		} else {
			return r + 14
		}
	} else if r >= 'A' && r <= 'Z' {

		if r >= 'M' {
			return r - 14
		} else {
			return r + 14
		}
	}

	return string(r)

}
