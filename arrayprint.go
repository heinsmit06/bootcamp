package bootcamp

import (
	"github.com/alem-platform/ap"
)

func ArrayPrint(arr [20]int) {
	for i := 0; i < len(arr); i++ {
		PutNumber(arr[i])

		if i < len(arr)-1 {
			ap.PutRune(' ')
		}
	}
}

func PutNumber(n int) {
	if n < 0 {
		ap.PutRune('-')
		n = -n
	}
	Recursion(n)
}

func Recursion(n int) {
	if n < 10 {
		ap.PutRune(rune(n + 48))
		return
	}
	Recursion(n / 10)
	ap.PutRune(rune(n%10 + 48))
}
