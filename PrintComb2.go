package piscine

import "fmt"

func PrintComb2() {
	var a, b, c, d int

	for a = 0; a < 10; a++ {
		for b = 0; b < 10; b++ {
			for c = 0; c < 10; c++ {
				for d = b + 1; d < 10; d++ {

					fmt.Print(a)
					fmt.Print(b)
					fmt.Print(" ")
					fmt.Print(c)
					fmt.Print(d)
					fmt.Print(",")
					fmt.Print(" ")

					if a == 9 && b == 7 && c == 9 && d == 9 {

						fmt.Print(9)
						fmt.Print(8)
						fmt.Print(" ")
						fmt.Print(9)
						fmt.Print(9)
						break

					}

				}
				if a == 9 && b == 7 && c == 9 && d == 9 {
					break
				}
			}
			if a == 9 && b == 7 && c == 9 && d == 9 {
				break
			}
		}
		if a == 9 && b == 7 && c == 9 && d == 9 {
			break
		}
	}

}
