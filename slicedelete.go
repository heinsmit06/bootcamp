package main

import "fmt"

func SliceDelete(arr *[]int, low, high int) bool {
	if (low < 0) || (high > len(*arr)) || !(low < high) || (*arr == nil) {
		return false
	}

	arr_len := len(*arr)
	arr_copy := *arr
	copy((*arr)[low:], arr_copy[high:])
	(*arr) = (*arr)[:(arr_len - (high - low))]
	return true
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5}

	fmt.Println(arr)                     // [0, 1, 2, 3, 4, 5]
	fmt.Println(SliceDelete(&arr, 1, 3)) // true
	fmt.Println(arr)                     // [0, 3, 4, 5]

	fmt.Println(SliceDelete(&arr, 0, 2)) // false
	fmt.Println(arr)                     // [0, 3, 4, 5]

	fmt.Println(SliceDelete(&arr, 5, 3)) // false
	fmt.Println(arr)                     // [0, 3, 4, 5]

	fmt.Println(SliceDelete(&arr, -10, 5)) // false
	fmt.Println(arr)                       // [0, 3, 4, 5]

	fmt.Println(SliceDelete(&arr, 0, 2)) // true
	fmt.Println(arr)                     // []
}
