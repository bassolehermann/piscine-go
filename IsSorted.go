package piscine

import "sort"

func IsSorted(f func(a, b int) int, tab []int) bool {

	if sort.IntsAreSorted(tab) {

		return true
	} else {
		return false
	}

}

func f(nbr1, nbr2 int) {

}
