package piscine



func IsSorted(f func(a, b int) int, tab []int) bool {
	
var booleen bool 	


		for i:=0;i<len(tab)-1;i++{

			if tab[i]>tab[i+1] {
				booleen=false
			}

		}


if booleen{
	return true
}else{
	return false
}


}


func f(nbr1,nbr2 int) int {

if nbr1>nbr2{
	return nbr1
}else if nbr1==nbr2{

	return 0
}else {

	return -nbr1
}




}