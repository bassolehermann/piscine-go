package piscine

import "fmt"

func PrintComb2() {
	var a, b, c, d int

	for a = 0; a < 10; a++ {
		for b = 0; b < 10; b++ {
			for c = a; c < 10; c++ {
				for d = b + 1; d < 10; d++ {
					fmt.Print(a)
					fmt.Print(b)
					fmt.Print(" ")
					fmt.Print(c)
					fmt.Print(d)
					fmt.Print(",")
					fmt.Print(" ")

				}
			}
		}
	}

}
