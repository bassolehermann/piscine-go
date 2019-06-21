package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {

	File,err := os.Open("quest8.txt")

	if os.Args == nil && err != nil  {

		fmt.Println("")
	} else if err != nil {
		fmt.Println(err.Error())
	} else if os.Args[1]=="quest8.txt" || os.Args[1]=="quest8T.txt" {

		result, _ := ioutil.ReadAll(File)
		fmt.Println(string(result))

	}
	

}