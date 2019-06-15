
package piscine

import (

"math"
)

func RecursiveFactorial(nb int) int {


var factVal =1


	if nb == 1 {

		return 1
	}
	
	factVal=nb*RecursiveFactorial(nb-1)
	
	if  factVal > math.MaxInt32 {

		return 0

	} else {
		
		return	factVal
	}

} 


