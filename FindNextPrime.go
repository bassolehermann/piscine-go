package piscine

import "math"

func FindNextPrime(nb int) int {

	if nb == 2 || nb == 3 {

		return nb
	}
	if nb<0{
		return 2
	}

	for 1000000093 > nb{

		if nb%1000000093 == 0 {
			
		}else if nb%1000000093 != 0 && float64(int(math.Sqrt(float64(nb)))) != math.Sqrt(float64(nb)) {
			return nb
			break
		}
		nb++
	}

return nb
}
