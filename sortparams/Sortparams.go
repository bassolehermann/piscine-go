package main


import (
	"fmt"
	"os"
)



func main() {
var temp byte 

		for i := 1; i < len (os.Args); i++ {

			for j := i+1 ; j < len(os.Args)+1; j++ {

				if os.Args[j]>os.Args[i]{

					temp=os.Args[j]
					os.Args[j]=os.Args[i]
					os.Args[i]=temp

				}
			}
		}

		for i := 1; i < len(os.Args); i++ {

			fmt.Printf("%v",os.Args[i])
			fmt.Printf("\n")
			
		}
	
}