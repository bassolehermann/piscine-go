package piscine

import "math"

func Sqrt(nb int) int {
	if nb%2 != 0 {
		return 0

	}

	nbr := math.Sqrt(float64(nb))
	return int(nbr)
}
