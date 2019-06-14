
package piscine


func RecursiveFactorial(nb int) int {

var factVal = 1
const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
const (
    MaxInt  = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
    MinInt  = -MaxInt - 1         // -1 << 31 or -1 << 63
    MaxUint = 1<<UintSize - 1     // 1<<32 - 1 or 1<<64 - 1
)

if nb <0 {

	return 0
}

if nb > 1{

	factVal=nb*RecursiveFactorial(nb-1)

	if factVal > MaxInt  {

		return 0
		}

	}
return factVal

}


