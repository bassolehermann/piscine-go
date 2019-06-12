package piscine

import "fmt"

func Putstr(str string) {

	for i := 0; i < len(str); i++ {
		fmt.Printf(string(str[i]))
		fmt.Print("%")
	}

}
