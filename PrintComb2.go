package piscine

import "fmt"

func PrintComb2() {
	var a, b, c, d int

	for a = 0; a < 10; a++ {
		for b = a + 1; b < 10; b++ {
			if a == 98 && b == 99 {

				fmt.Print(a)
				fmt.Print(" ")
				fmt.Print(b)
				fmt.Print("\n")

			} else {
				fmt.Print(a)
				fmt.Print(" ")
				fmt.Print(b)
				fmt.Print(",")
				fmt.Print(" ")
			}

		}
	}

}
