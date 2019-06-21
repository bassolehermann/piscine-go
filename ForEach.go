

package piscine

import "fmt"

func ForEach(f func(int)int, arr []int) []int {
var a = make([]int, len(arr), len(arr))
    for i, val := range arr{

        a[i] = f(val)
	fmt.Print(a[i])
	
    }
    return a
}

func PutNbr (nbr int) int {

	return nbr

}