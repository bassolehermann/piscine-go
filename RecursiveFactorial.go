
package piscine


func RecursiveFactorial(nb int) int {

const wordSize = 32 << (^uint(0) >> 63)
const MinInt = -(1 << (wordSize - 1))
const MaxInt = ^MinInt
const MinUint = 0
const MaxUint = MaxInt << 1 + 1	

var factVal int


	if nb == 1 {

		return 1
	}

	factVal=nb*RecursiveFactorial(nb-1)
	

	if  factVal > MaxInt{

		return 0

	} else {
		
		return	factVal
	}


}


