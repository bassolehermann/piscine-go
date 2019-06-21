package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {

	File,err := os.Open("quest8.txt")

	if len(os.Args) == 0 {

		fmt.Println(os.Args[0])
	} else if err != nil {
		fmt.Println(err.Error(os.Args[0]))
	} else if os.Args[1]=="quest8.txt" || os.Args[1]=="quest8T.txt" {

		result, _ := ioutil.ReadAll(File)
		fmt.Println(string(result))

	}
	

}