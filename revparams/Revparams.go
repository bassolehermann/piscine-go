package main


import (
	"fmt"
	"os"
)



func main() {

	
	for i, j := 0, len(os.Args)-1; i < j; i, j = i+1, j-1 {
		os.Args[i], os.Args[j] = os.Args[j], os.Args[i]
	}

	fmt.Println(os.Args)
}