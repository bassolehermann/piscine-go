package main


import (
	"fmt"
	"strings"
	"os"
)

func reverse_words(s ) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, "\n")
}

func main() {
	fmt.Println(reverse_words(os.Args))
}