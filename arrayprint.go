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
