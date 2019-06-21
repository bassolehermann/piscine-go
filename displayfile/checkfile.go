package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("quest8.txt")

	if len(os.Args) >= 2 {
		fmt.Println("Too many arguments")
	} else if err != nil {
		fmt.Println("File name missing")
	} else {
		arg := make([]byte, 31)
		file.Read(arg)
		fmt.Println(string(arg))
		file.Close()
	}

}
