package bootcamp

import (
	"github.com/alem-platform/ap"
)

func PutDigit(n int) {
	if n >= 0 && n < 10 {
		ap.PutRune(rune(n + 48))
	}
}
