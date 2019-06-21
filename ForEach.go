

package piscine



func ForEach(f func(int)int, arr []int) []int {
var a = make([]int, len(arr), len(arr))
    for index, val := range arr{

        a[index] = f(val)
    }
    return a
}

func PutNbr (nbr int) int {

	return nbr

}