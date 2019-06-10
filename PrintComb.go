package piscine

import "fmt"

func PrintComb() {

	var n1 = '0'
	var n2 = '0'
	var n3 = '0'

	for n1 <= '9' {
		for n2 <= '9' {
			for n3 <= '9' {
				if n1 <= n2 && n2 <= n3 {
					fmt.Println(0, &n1, 1)
					fmt.Println(0, &n2, 1)
					fmt.Println(0, &n3, 1)
					fmt.Println(0, "\n", 1)
				}
				n3 = n3 + 1
			}
			n3 = '0'
			n2 = n2 + 1
		}
		n2 = '0'
		n1 = n1 + 1
	}

}
