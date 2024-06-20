package main

import "github.com/alem-platform/ap"

func main() {
	s := "Salem!"

	for i := 0; i < len(s); i++ {
		ap.PutRune(rune(s[i]))
	}

	ap.PutRune('\n')
}
