package piscine

import "fmt"

func PrintStr(str string) {

	for _, word := range str {

		fmt.Printf(string(word))
	}

}
