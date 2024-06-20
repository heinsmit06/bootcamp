package main

import (
	"bootcamp"
	"fmt"

	"github.com/alem-platform/ap"
)

func main() {
	var int1 int
	var int2 int
	fmt.Scanf("%d %d", &int1, &int2)

	bootcamp.PutNumber(int1 + int2)
	ap.PutRune(' ')
	bootcamp.PutNumber(int1 - int2)
	ap.PutRune(' ')
	bootcamp.PutNumber(int1 * int2)
	ap.PutRune(' ')

	if int2 != 0 {
		bootcamp.PutNumber(int1 / int2)
	} else {
		ap.PutRune('F')
	}
}
