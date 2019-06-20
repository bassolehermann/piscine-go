package main

import "fmt"
import "os"

func printStr(str string) {
	
	fmt.Println(str)
}

func isEven(nbr int) bool {
	if nbr == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	if isEven(len(os.Args)){
		printStr("I have an odd number of arguments")
	} else {
		printStr("I have an even number of arguments")
	}
}