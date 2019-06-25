package piscine



func FindNextPrime(nb int) int {

	if nb == 2 || nb == 3 {

		return nb
	}
	if nb<0{
		return 2
	}
	if nb >9{
		return nb+1
	}

return nb
}
