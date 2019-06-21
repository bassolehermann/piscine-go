

package piscine



func ForEach(f func(int)int, arr []int) []int {
var a = make([]int, len(arr), len(arr))
    for i, val := range arr{

        a[i] = f(val)
    }
    return a
}

func PutNbr (nbr int) int {

	return nbr

}