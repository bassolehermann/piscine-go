package main

import (
	"fmt"
	"os"
)

func main() {

	file, _ := os.Open("quest8.txt")

	if os.Args == nil {
		fmt.Println("File name missing")
	} else if len(os.Args) >= 1 {
		fmt.Println("Too many arguments")
	} else {
		arg := make([]byte, 31)
		file.Read(arg)
		fmt.Println(string(arg))
		file.Close()
	}

}
