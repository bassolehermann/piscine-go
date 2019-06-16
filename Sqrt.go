package piscine

import "math"

func Sqrt(nb int) int {
	if nb < 0 {

		return 0

	}

	if float64(int(math.Sqrt(float64(nb)))) == math.Sqrt(float64(nb)) {

		return int(math.Sqrt(float64(nb)))

	} else {
		return 0

	}
}
