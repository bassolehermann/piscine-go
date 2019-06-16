package piscine

import "math"

func IsPrime(nb int) bool {

	if nb%2 == 0 || nb%3 == 0 {

		return false

	}
	if float64(int(math.Sqrt(float64(nb)))) == math.Sqrt(float64(nb)) {

		return false

	}
	for i := 3; i < int(math.Sqrt(float64(nb))); i = i + 2 {

		if nb%i == 0 {
			return false
		}

	}

	return true
}
