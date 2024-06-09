package main

import (
	"bootcamp"
	"fmt"
)

func main() {
	arr := [20]int{1, -2, 3, 4, 5, -6, 0, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	bootcamp.ArrayPrint(arr)
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println("input is:", n)
}
