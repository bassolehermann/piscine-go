package piscine


func Abort(a, b, c, d, e int) int {
var result int

	x:=[]int={a,b,c,d,e}

	int.Sort(x)

	result = x[1] / x[3]

	return result

}