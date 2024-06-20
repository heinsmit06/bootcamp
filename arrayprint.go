package bootcamp

import (
	"github.com/alem-platform/ap"
)

func ArrayPrint(arr [20]int) {
	for i := 0; i < len(arr); i++ {
		PutNumber1(arr[i])

		if i < len(arr)-1 {
			ap.PutRune(' ')
		}
	}
}

func PutNumber1(n int) {
	if n < 0 {
		ap.PutRune('-')
		n = -n
	}
	Recursion1(n)
}

func Recursion1(n int) {
	if n < 10 {
		ap.PutRune(rune(n + 48))
		return
	}
	Recursion1(n / 10)
	ap.PutRune(rune(n%10 + 48))
}
