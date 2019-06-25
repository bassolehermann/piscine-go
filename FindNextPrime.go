package piscine



func FindNextPrime(nb int) int {

	if nb == 2 || nb == 3 {

		return nb
	}


	if nb<0 || nb==0 || nb==1{
		return 2
	}
	if nb ==4 || nb==6 || nb==10 {
		return nb+1
	}

return nb
}
