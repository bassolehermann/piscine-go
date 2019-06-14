
package piscine


func RecursiveFactorial(nb int ) int { 
var fact =1
const wordSize = 32 << (^uint(0) >> 63)
const MinInt = -(1 << (wordSize - 1))
const MaxInt = ^MinInt
const MinUint = 0
const MaxUint = MaxInt << 1 + 1	
	
	if nb==1 {
	
	return 1
	}
	
	fact=nb * RecursiveFactorial(nb-1)
	
	if nb<0 || fact > MaxInt{
	
	return 0
	
	}else{
	
		return fact
	}
}


