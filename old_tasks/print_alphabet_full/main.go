package main

import "github.com/alem-platform/ap"

func main() {
	for i := 97; i <= 122; i++ {
		ap.PutRune(rune(i))
		ap.PutRune(rune(i - 32))
	}

	ap.PutRune('\n')
}
