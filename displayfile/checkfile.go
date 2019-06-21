package main

import (
	"fmt"
	"os"
)

func main() {

	//file, err := os.Open("quest8.txt")

	if os.Args == nil {
		fmt.Println(err)
	} else if len(os.Args) == 1 {
		
		fmt.Println("file exist")
	} else if len(os.Args) >1 {
		fmt.Println("Too many arguments")
	}

}
