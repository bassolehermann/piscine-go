package main


import (
	"fmt"
	"os"
)



func main() {

	
	for i, j := 1, len(os.Args)-1; i < j; i, j = i+1, j-1 {
		os.Args[i], os.Args[j] = os.Args[j], os.Args[i]
	}

	for i:=1;i<len(os.Args);i++{


		fmt.Printf("%v",os.Args[i])
		fmt.Printf("\n")

	}

	
}