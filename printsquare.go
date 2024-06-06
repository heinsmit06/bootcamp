package bootcamp

import "github.com/alem-platform/ap"

func PrintSquare(w int) {
	for r := 0; r < w; r++ {
		for c := 0; c < w; c++ {
			ap.PutRune('0')

			if c == w-1 {
				break
			} else {
				ap.PutRune(' ')
			}

		}

		//if r == w-1 {
		//	break
		//} else {
		ap.PutRune('\n')
		//}

	}
}
