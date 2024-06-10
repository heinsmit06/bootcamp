package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

func main() {
	ap.PutRune(0)

	slc := [4]int{}
	slc[2] = 165
	fmt.Println(slc)

	arr := make([]int, 5)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	fmt.Println(arr[3:])
	fmt.Println(arr[:2])
}
