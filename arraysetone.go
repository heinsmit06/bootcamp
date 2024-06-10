package bootcamp

func ArraySetOne(arr *[20]int, idx int, val int) bool {
	if (arr == nil) || (idx >= 20) {
		return false
	}

	arr[idx] = val
	return true
}
