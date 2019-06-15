package piscine

func IterativePower(nb int, power int ) int{ 

		
        	for i:= 1; i<=power; i++ {
            		nb=nb*i   
        		}
		if power <0 {
        		return 0
        	}else{
	
				return nb 
        	} 
}