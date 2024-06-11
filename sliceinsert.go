package bootcamp

func SliceInsert(arr *[]int, idx int, values ...int) bool {
	//	if idx < 0 {
	//		return false
	//	}
	//
	//	arr_len := len(*arr)
	//	slc := []int{}
	//	for i := idx; i < arr_len; i++ {
	//		slc = append(slc, (*arr)[i])
	//	}
	//
	//	checker_flag := 0
	//	if idx >= arr_len {
	//		for _, v := range values {
	//			*arr = append(*arr, v)
	//		}
	//
	//		checker_flag = 1
	//	}
	//
	//	values_idx_counter := 0
	//
	//	if checker_flag != 1 {
	//		for j := idx; j < arr_len+len(values)-1; j++ {
	//
	//			if idx >= arr_len {
	//				*arr = append((*arr), values[values_idx_counter])
	//				values_idx_counter++
	//				continue
	//			}
	//
	//			(*arr)[j] = values[values_idx_counter]
	//			values_idx_counter++
	//		}
	//	}
	//
	//	for k := 0; k < len(slc); k++ {
	//		(*arr) = append((*arr), slc[k])
	//	}
	//
	//	return true

	if idx < 0 || idx > len(*arr) {
		return false
	}

	*arr = append(*arr, values...)
	copy((*arr)[idx+len(values):], (*arr)[idx:])
	copy((*arr)[idx:], values)

	return true
}
