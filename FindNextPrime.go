package piscine

import "math"

func FindNextPrime(nb int) int {

	if nb == 2 || nb == 3 {

		return nb
	}

	if nb%2 == 0 || nb%3 == 0 {

		return nb + 1

	}
	if float64(int(math.Sqrt(float64(nb)))) == math.Sqrt(float64(nb)) {

		return nb + 1

	}
	for i := 3; i < int(math.Sqrt(float64(nb))); i = i + 2 {

		if nb%i == 0 {
			return nb + 1
		}

	}

	return nb
}
