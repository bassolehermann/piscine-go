package piscine

package "math"
func IterativeFactorial(nb int) int { 

var factVal = 1 
			
			
        	if nb < 0 || math.Maxint  {
        		return 0
        	}else{

        	for i:= 1; i<=nb; i++ {
            factVal *= (i)  
        	}
				return factVal 
        	} 
}