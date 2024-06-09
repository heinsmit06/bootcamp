package main

import (
	"bootcamp"
	"fmt"

	"github.com/alem-platform/ap"
)

func main() {
	var string_length int
	var str string
	fmt.Scanf("%d", &string_length)
	fmt.Scanf("%s", &str)

	for i := 0; i < string_length; i++ {
		bootcamp.PutNumber(int(rune(str[i])))

		if i < string_length-1 {
			ap.PutRune(' ')
		}
	}
}
