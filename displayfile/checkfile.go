package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("quest8.txt")

	if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else if len(os.Args) == 2 {
		
		fmt.Println(err.Error())

	} else if len(os.Args) >2 {
		fmt.Println("Too many arguments")
	} 

}
