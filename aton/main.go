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
	string_array := make([]string, 1)

	fmt.Scanf("%s", &string_array[0])

	for i := 0; i < string_length; i++ {

		bootcamp.PutNumber(int(rune(string_array[0][i])))

		if i < len(str)-1 {
			ap.PutRune(' ')
		}
	}
}
