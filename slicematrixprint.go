package bootcamp

import "github.com/alem-platform/ap"

func SliceMatrixPrint(matrix [][]rune) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			ap.PutRune(matrix[i][j])

			if j < len(matrix[0])-1 {
				ap.PutRune(' ')
			}
		}
		ap.PutRune('\n')
	}
}
