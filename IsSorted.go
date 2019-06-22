package piscine

import "sort"

func IsSorted(f func(a, b int) int, tab []int) bool {
	var s []int	

		s=sort.Ints(tab)
		if tab==s {
			return true
		}else {
			return false
		}


}


func f(nbr1,nbr2 int) int {

	if nbr1> nbr2{

		return nbr1
	} else if nbr1==nbr2{

		return 0
	}else {

	
		return -nbr1
	}
}