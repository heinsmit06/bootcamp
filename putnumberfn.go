package bootcamp

import (
	"github.com/alem-platform/ap"
)

func PutNumber(n int) {
	//if n >= 0 && n < 10 {
	//	ap.PutRune(rune(n + 48))
	//} else if n >= 10 && n < 100 {
	//	ap.PutRune(rune(n/10 + 48))
	//	ap.PutRune(rune(n%10 + 48))
	//} else if n >= 100 && n < 1000 {
	//	ap.PutRune(rune(n/100 + 48))
	//	n = n - (n/100)*100
	//	ap.PutRune(rune(n/10 + 48))
	//	ap.PutRune(rune(n%10 + 48))
	//} else if n >= 1000 && n < 10000 {
	//	ap.PutRune(rune(n/1000 + 48))
	//	n = n - (n/1000)*1000
	//	ap.PutRune(rune(n/100 + 48))
	//	n = n - (n/100)*100
	//	ap.PutRune(rune(n/10 + 48))
	//	ap.PutRune(rune(n%10 + 48))
	//}
	//
	//if n < 0 && n > -10 {
	//	ap.PutRune('-')
	//	ap.PutRune(rune((n)*(-1) + 48))
	//} else if n <= -10 && n > -100 {
	//	ap.PutRune('-')
	//	ap.PutRune(rune((n/10)*(-1) + 48))
	//	ap.PutRune(rune((n%10)*(-1) + 48))
	//} else if n <= -100 && n > -1000 {
	//	ap.PutRune('-')
	//	ap.PutRune(rune((n/100)*(-1) + 48))
	//	n = n - (n/100)*100
	//	ap.PutRune(rune((n/10)*(-1) + 48))
	//	ap.PutRune(rune((n%10)*(-1) + 48))
	//} else if n <= -1000 && n > -10000 {
	//	ap.PutRune('-')
	//	ap.PutRune(rune((n/1000)*(-1) + 48))
	//	n = n - (n/1000)*1000
	//	ap.PutRune(rune((n/100)*(-1) + 48))
	//	n = n - (n/100)*100
	//	ap.PutRune(rune((n/10)*(-1) + 48))
	//	ap.PutRune(rune((n%10)*(-1) + 48))
	//}
	//
	//if (n < 0) {
	//	ap.PutRune('-')
	//	n = -n
	//}
	//Recursion(n)
}

func Recursion(n int) {
	if n < 10 {
		ap.PutRune(rune(n + 48))
		return
	}
	Recursion(n / 10)
	ap.PutRune(rune(n%10 + 48))
}
