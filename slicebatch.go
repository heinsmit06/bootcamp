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
		if k == len(partitioned_slice)-1 {
			partitioned_slice[k] = make([]int, len(slice)%size)
			continue
		}
		partitioned_slice[k] = make([]int, size)
	}

	for i := 0; i < len(partitioned_slice); i++ {
		for j := 0; j < len(partitioned_slice[0]); j++ {
			if idx_counter == len(slice) {
				break
			}

			partitioned_slice[i][j] = slice[idx_counter]
			idx_counter++
		}
	}

	return partitioned_slice
}
