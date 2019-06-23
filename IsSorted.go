package piscine

import "sort"

func IsSorted(f func(a, b int) int, tab []int) bool {
	var booleen bool

	if sort.IntsAreSorted(tab) {

		booleen = true
	} else if sort.IntsAreSorted(reverseInts(tab)) {
		booleen = true
	} else {
		booleen = false
	}

	if booleen {
		return true
	} else {
		return false
	}

}

func reverseInts(tab []int) []int {
	if len(tab) == 0 {
		return tab
	}
	return append(reverseInts(tab[1:]), tab[0])
}

func f(nbr1, nbr2 int) {

}
