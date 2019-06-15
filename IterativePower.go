package piscine

import "math"


func IterativePower(nb int, power int ) int{ 

var factVal int 

        	
		
        	for i:= 1; i<=nb; i++ {
            		factVal= math.Pow(i,power)  
        		}
		if nb < 0 && power <0 {
        		return 0
        	}else{
	
				return factVal 
        	} 
}