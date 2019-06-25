package piscine

import "math"

func FindNextPrime(nb int) int {

	if nb == 2 || nb == 3 {

		return nb
	}

	if nb%2 == 0 || nb%3 == 0 {

		

	}
	if float64(int(math.Sqrt(float64(nb)))) == math.Sqrt(float64(nb)) {

	

	}
	for i := 3; i < int(math.Sqrt(float64(nb))); i = i + 2 {

		if nb%i == 0 {
			
		}else {
			return i
			break
		}

	}

	return nb
}
