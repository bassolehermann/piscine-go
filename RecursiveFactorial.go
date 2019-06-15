
package piscine

import (

"math"
)

func RecursiveFactorial(nb int) int {


var factVal int


	if nb == 1 {

		return 1
	}
	
	
	
	if  factVal > math.MaxInt32 {

		return 0

	} else {
		factVal=nb*RecursiveFactorial(nb-1)
		return	factVal
	}
	

} 


