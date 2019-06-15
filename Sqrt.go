package piscine

import "math"

func Sqrt(nb int) int {
	if int(math.Sqrt(float64(nb)))%2 != 0 {
		return 0

	}

	return int(math.Sqrt(float64(nb)))
}
