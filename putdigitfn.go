package bootcamp

import "github.com/alem-platform/ap"

func PutDigit(n int) {
	ap.PutRune(rune(n + 48))
}
