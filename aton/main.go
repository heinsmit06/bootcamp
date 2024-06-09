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

		if str[i] == '\t' {
			bootcamp.PutNumber(int(rune(str[i])))
		} else if str[i] == '\n' {
			bootcamp.PutNumber(int(rune(str[i])))
		} else if str[i] == '\v' {
			bootcamp.PutNumber(int(rune(str[i])))
		} else if str[i] == '\f' {
			bootcamp.PutNumber(int(rune(str[i])))
		} else if str[i] == '\r' {
			bootcamp.PutNumber(int(rune(str[i])))
		} else if str[i] == ' ' {
			bootcamp.PutNumber(int(rune(str[i])))
		} else if str[i] == 0x85 {
			bootcamp.PutNumber(int(rune(str[i])))
		} else if str[i] == 0xA0 {
			bootcamp.PutNumber(int(rune(str[i])))
		} else {
			bootcamp.PutNumber(int(rune(str[i])))
		}

		if i < len(str)-1 {
			ap.PutRune(' ')
		}
	}
}
