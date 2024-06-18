package bootcamp

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
