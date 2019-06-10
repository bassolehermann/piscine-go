package main

import piscine "fmt"


func IsNegative(nb int) {

	if nb<0 {
		piscine.Println("T")
	} else {
		piscine.Println("F")
	}
}

func main() {

	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}