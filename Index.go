package piscine

func Index(s string, toFind string) int {
	var count int
	count = 0
	for _, word := range s {
		count++
		if string(word) == string(toFind[0]) {

			return count - 1
		}
	}
	return count - 1
}
