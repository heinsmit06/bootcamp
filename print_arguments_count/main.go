package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	args := os.Args

	numArgs := len(args) - 1

	var numArgsStr string
	if numArgs == 0 {
		numArgsStr = "0"
	} else {
		numArgsStr = ""
		for numArgs > 0 {

			digit := byte('0' + numArgs%10)

			numArgsStr = string(digit) + numArgsStr

			numArgs /= 10
		}
	}

	for _, char := range numArgsStr {
		ap.PutRune(char)
	}

	ap.PutRune('\n')
}
