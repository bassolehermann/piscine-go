package main 

import piscine "fmt"

func IsNegative(nb int) {

	if nb<0 {
		fmt.Println("T")
	} else {
		fmt.Println("F")
	}
}

func main() {

	piscine.IsNegative(1)
	piscine.IsNegative(0)
	piscine.IsNegative(-1)
}