
package piscine


func RecursiveFactorial(nb int ) int { 
const wordSize = 32 << (^uint(0) >> 63)
const MinInt = -(1 << (wordSize - 1))
const MaxInt = ^MinInt
const MinUint = 0
const MaxUint = MaxInt << 1 + 1	
	
	
	if nb == 0 {
        return 1
    }
	if (nb * RecursiveFactorial(nb-1))> MaxInt || nb<0 {
	
	return 0
	
	}
    
	return nb * RecursiveFactorial(nb-1)
	
}


