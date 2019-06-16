package piscine

func Index(s string, toFind string) int {
	var i int

	for i = 0; i < len(s); i++ {
		if string(s[i]) == string(toFind[0]) {
			return (i)
		}
	}
	return -1
}
