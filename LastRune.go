package piscine

func LastRune(s string) rune {

	x := len(s)

	r := []rune(s)

	return r[x-1]

}
