package piscine

import "math"

func FindNextPrime(nb int) int {

	if nb == 2 || nb == 3 {

		return nb
	}
	if nb<0{
		return 2
	}
	if nb<7{
		return 7
	}
	if nb<11{
		return 11
	}
	if nb<13{
		return 13
	}
	if float64(int(math.Sqrt(float64(nb)))) == math.Sqrt(float64(nb)) {

	

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
