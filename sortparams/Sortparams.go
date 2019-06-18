package main


import (
	"fmt"
	"os"
)



func main() {
var temp byte 
var arg byte
arg=[]byte(os.Args)


		for i := 1; i < len (os.Args); i++ {

			for j := i+1 ; j < len(os.Args)+1; j++ {

				if arg[j]>arg[i]{

					temp=arg[j]
					arg[j]=arg[i]
					arg[i]=temp

				}
			}
		}

		for i := 1; i < len(os.Args); i++ {

			fmt.Printf("%v",os.Args[i])
			fmt.Printf("\n")
			
		}
	
}