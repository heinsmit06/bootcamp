package bootcamp

func RunesLen(arr [20]rune) int {
	arr_len := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			break
		}

		arr_len++
	}

	return arr_len
}
