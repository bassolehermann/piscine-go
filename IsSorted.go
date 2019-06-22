package piscine

import "sort"

func IsSorted(f func(a, b int) int, tab []int) bool {
	
var temp int
var t []int 	


		for i:=0;i<len(tab);i++{

			for i:= 0 ;i<len(tab)-1;j++{

				if tab[i]>tab[i+1] {
					temp=tab[i]
					 tab[i]=tab[i+1]
					 tab[i+1]=temp
				}
			}
			t= t[i]
		}



		if tab==t {
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