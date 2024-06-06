package main

import "github.com/alem-platform/ap"

func main() {
	for h := 0; h < 24; h++ {
		for m := 0; m < 60; m++ {

			if h < 10 {
				ap.PutRune(48)
			} else if h >= 10 && h < 20 {
				ap.PutRune(49)
			} else {
				ap.PutRune(50)
			}

			ap.PutRune(rune(h%10 + 48))
			ap.PutRune(':')

			if m < 10 {
				ap.PutRune(48)
			} else if m >= 10 && m < 20 {
				ap.PutRune(49)
			} else if m >= 20 && m < 30 {
				ap.PutRune(50)
			} else if m >= 30 && m < 40 {
				ap.PutRune(51)
			} else if m >= 40 && m < 50 {
				ap.PutRune(52)
			} else {
				ap.PutRune(53)
			}

			ap.PutRune(rune(m%10 + 48))

			ap.PutRune('\n')

		}
	}
}
