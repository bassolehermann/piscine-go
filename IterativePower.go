package piscine

import "math"

func IterativePower(nb int, power int) int {
	var result int

	c := math.Pow(2, 0)

	if power < 0 {
		return 0
	} else {
		result = int(c)
		return result
	}
}
