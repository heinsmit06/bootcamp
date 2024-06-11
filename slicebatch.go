package bootcamp

func SliceBatch(slice []int, size int) [][]int {
	slc_size := 0

	if (len(slice))%size == 0 {
		slc_size = len(slice) / size
	} else {
		slc_size = (len(slice) / size) + 1
	}

	idx_counter := 0

	partitioned_slice := make([][]int, slc_size)
	for k := 0; k < len(partitioned_slice); k++ {
		partitioned_slice[k] = make([]int, size)
	}

	for i := 0; i < slc_size; i++ {
		for j := 0; j < size; j++ {
			if idx_counter != len(slice) {
				partitioned_slice[i][j] = slice[idx_counter]
				idx_counter++
			}
		}
	}

	return partitioned_slice
}
