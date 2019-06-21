package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("quest8.txt")

	if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else if len(os.Args) == 2 {
		
		fmt.Println(err.Erro())

	} else if len(os.Args) >2 {
		fmt.Println("Too many arguments")
	}
}
