package main

import (
	"github.com/alem-platform/ap"
)

func main() {
	for i := 0; i <= 9; i++ {
		ap.PutRune(rune(57 - i))
	}

	ap.PutRune('\n')
}
