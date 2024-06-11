package bootcamp

func SliceBatch(slice []int, size int) [][]int {
	if size <= 0 {
		return nil
	}

	partitioned_slice := [][]int{}

	for i := 0; i < len(slice); i += size {
		batch_end := i + size

		if batch_end > len(slice) {
			batch_end = len(slice)
		}

		partitioned_slice = append(partitioned_slice, slice[i:batch_end])
	}

	return partitioned_slice
}
