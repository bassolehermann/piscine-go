package piscine

import"sort"
func Abort(a, b, c, d, e int) int {
var result int

	x:=[]int{a,b,c,d,e}

	sort.Ints(x)

	result = x[2]

	return result

}