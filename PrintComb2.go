package piscine

import "fmt"

func PrintComb2() {
	var a, b int

	for a = 0; a <= 99; a++ {
		for b = a + 1; b <= 99; b++ {
			if a == 98 && b == 99 {

				fmt.Printf("%02d", a)
				fmt.Print(" ")
				fmt.Printf("%02d", b)
				fmt.Print("\n")

			} else {
				fmt.Printf("%02d", a)
				fmt.Print(" ")
				fmt.Printf("%02d", b)
				fmt.Print(",")
				fmt.Print(" ")
			}

		}
	}

}
