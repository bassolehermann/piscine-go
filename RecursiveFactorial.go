
package piscine


func RecursiveFactorial(nb int ) int {
var factVal = 1 
const wordSize = 32 << (^uint(0) >> 63)
const MinInt = -(1 << (wordSize - 1))
const MaxInt = ^MinInt
const MinUint = 0
const MaxUint = MaxInt << 1 + 1	

	factVal = nb * RecursiveFactorial(nb-1)

	if nb == 0 || factVal > MaxInt {
        return 0
    }else {
    	return factVal	
    }
    
}


