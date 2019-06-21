package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("quest8.txt")

	if os.Args == nil {
		fmt.Println("File name missing")
	} else if len(os.Args) >= 1 {
		fmt.Println(err)
	} else {
		arg := make([]byte, 31)
		file.Read(arg)
		fmt.Println(string(arg))
		file.Close()
	}

}
