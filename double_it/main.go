package main

import (
	"bootcamp"
	"fmt"

	"github.com/alem-platform/ap"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	arr := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	for j := 0; j < N; j++ {
		bootcamp.PutNumber(arr[j] * 2)
		ap.PutRune(' ')
	}

	ap.PutRune('\n')
}
