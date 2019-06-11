package piscine

import "fmt"

func PrintComb() {

	var a int
	var b int
	var c int

	a = 0
	for a = 0; a < 10; a++ {
		for b = a + 1; b < 10; b++ {
			for c = b + 1; c < 10; c++ {

				fmt.Print(a, b, c, ",", " ")

			}

		}

	}

}
