package main

import "fmt"
import "os"

func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		fmt.Println(arrayStr[i])
	}
	fmt.Println('\n')
}

func isEven(nbr string) bool {
	if int(nbr) == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	if isEven(os.Args) == 1 {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}