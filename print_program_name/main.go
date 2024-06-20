package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	args := os.Args[0]

	var n int
	for i := len(args) - 1; i >= 0; i-- {
		if args[i] == '.' || args[i] == '/' {
			n = i + 1
			break
		}
	}
	str := args[n:]

	for _, char := range str {
		ap.PutRune(char)
	}
	ap.PutRune('\n')
}
