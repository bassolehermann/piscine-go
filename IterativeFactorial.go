package piscine


func IterativeFactorial(nb int) int { 

var factVal = 1 
			for i:= 1; i<=nb; i++ {
            factVal *= (i)  
        	}

        	if nb < 0  {
        		return 0
        	}else{
				return factVal 
        	} 
}