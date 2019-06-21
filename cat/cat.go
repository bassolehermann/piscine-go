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
		fmt.Println(err.Error())
	} else if os.Args==quest8.txt || os.Args==quest8T.txt {

		result, _ := ioutil.ReadAll(File)
		fmt.Println(string(result))

	}
	

}