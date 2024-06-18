package bootcamp

import "github.com/alem-platform/ap"

func PrintAlphabetLn() {
	for i := 97; i <= 122; i++ {
		ap.PutRune(rune(i))
	}

	ap.PutRune('\n')
}
