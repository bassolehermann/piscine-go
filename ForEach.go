

package piscine

import "fmt"


func ForEach(f func(int), arr []int) {

    for i, _:= range arr{
	i=i+1
	fmt.Print(i)
	
    }  
}


func PutNbr (arr int) { 
	
}