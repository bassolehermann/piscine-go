package main

import (
	"fmt"
	"os"
)

func main() {

	file, _ := os.Open("quest8.txt")

	if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else if len(os.Args) == 2 {
		
		tab:=make([]byte ,14)
		file.Read(tab)
		fmt.Println(string(tab))

	} else if len(os.Args) >2 {
		fmt.Println("Too many arguments")
	}

}
