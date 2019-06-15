package piscine

import (

"math"
)

func IterativePower(int nb, int power) int {

var result int 


	if power<0 {

		return 0
	}else{


	result= math.Pow(nb,power)	
	}

	

return result
}