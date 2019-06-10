package piscine

import "fmt"

func PrintComb() {

	var a int
	var b int
	var c int

	a = 0
	for a < b && b < c {
		for a <= 9 {
			b := a + 1
			for b <= 9 {
				c := b + 1
				for c <= 9 {
					fmt.Println("%d%d%d\n", a, b, c)
					c++
				}
				b++
			}
			a++
		}
	}

}
