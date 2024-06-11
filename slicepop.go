package bootcamp

func SlicePop(arr *[]int) int {
	arr_len := len(*arr)

	if arr_len == 0 {
		return 0
	}

	last_element := (*arr)[arr_len-1]
	(*arr) = (*arr)[:arr_len-1]
	return last_element
}
