package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var result int

	if len(os.Args) == 3 {

		if os.Args[1] == "+" {

			result = strconv.Atoi(os.Args[0]) + strconv.Atoi(os.Args[2])
			fmt.Print(result)

		} else if os.Args[1] == "-" {

			result = strconv.Atoi(os.Args[0]) - strconv.Atoi(os.Args[2])
			fmt.Print(result)

		} else if os.Args[1] == "*" {

			result = strconv.Atoi(os.Args[0]) * strconv.Atoi(os.Args[2])
			fmt.Print(result)

		} else if os.Args[1] == "/" {
			if strconv.Atoi(os.Args[2]) == 0 {

				fmt.Println("No division by 0")

			} else {
				result = strconv.Atoi(os.Args[0]) / strconv.Atoi(os.Args[2])
				fmt.Print(result)
			}
		} else if os.Args[1] == "%" {
			if strconv.Atoi(os.Args[2]) == 0 {

				fmt.Println("No modulo by 0")

			} else {
				result = strconv.Atoi(os.Args[0]) % strconv.Atoi(os.Args[2])
				fmt.Print(result)
			}
		} else {
			fmt.Print(0)
		}

	}
}
