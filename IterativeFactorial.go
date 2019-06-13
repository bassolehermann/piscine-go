package piscine

func IterativeFactorial(nb int) uint64 { 

var factVal uint64 = 1 // uint64 is the set of all unsigned 64-bit integers.
                       // Range: 0 through 18446744073709551615. 

  
    if(nb < 0){
        return 0   
    }else{        
        for i:=1; i<=nb; i++ {
            factVal *= uint64(i)  // mismatched types int64 and int
        }
         
    }    
    return factVal  /* return from function*/
}