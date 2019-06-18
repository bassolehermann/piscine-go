package main


import (
	"fmt"
	"os"
)

var temp int 

func main() {

		for i := 1; i < len (os.Args); i++ {

			for j := i+1 ; j < len(os.Args)+1; j++ {

				if int(os.Args[j])>int(os.Args[i]){

					temp=int(os.Args[j])
					int(os.Args[j])=int(os.Args[i])
					int(os.Args[i])=temp

				}
			}
		}

		for i := 1; i < len(os.Args); i++ {

			fmt.Printf("%v",os.Args[i])
			fmt.Printf("\n")
			
		}
	
}