package main

import (
	"fmt"
	"os"
)

func main() {

	//file, err := os.Open("quest8.txt")

	if len(os.Args) == 1 {
		fmt.Println("file name missing")
	} else if len(os.Args) == 2 {
		
		fmt.Println("file exist")
	} else if len(os.Args) >2 {
		fmt.Println("Too many arguments")
	}

}
