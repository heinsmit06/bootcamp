package bootcamp

func SliceInsert(arr *[]int, idx int, values ...int) bool {
	if idx < 0 {
		return false
	}

	arr_len := len(*arr)
	slc := []int{}
	for i := idx; i < arr_len; i++ {
		slc = append(slc, (*arr)[i])
	}

	if idx >= arr_len {
		for _, v := range values {
			*arr = append(*arr, v)
		}
	}

	values_idx_counter := 0

	for j := idx; j < arr_len; j++ {
		(*arr)[j] = values[values_idx_counter]
		values_idx_counter++
	}

	for k := 0; k < len(slc); k++ {
		(*arr) = append((*arr), slc[k])
	}

	return true
}
