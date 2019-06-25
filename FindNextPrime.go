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
	for 1000 > nb{

		if nb%1000 == 0 {
			
		}else if nb%1000 != 0 && float64(int(math.Sqrt(float64(nb)))) != math.Sqrt(float64(nb)) {
			return nb
			break
		}
		nb++
	}

return nb
}
