package main

import "github.com/alem-platform/ap"

func main() {
	for h := 0; h < 24; h++ {
		for m := 0; m < 60; m++ {

			ap.PutRune(rune(h/10 + 48))
			ap.PutRune(rune(h%10 + 48))
			ap.PutRune(':')
			ap.PutRune(rune(m/10 + 48))
			ap.PutRune(rune(m%10 + 48))

			ap.PutRune('\n')

		}
	}
}
