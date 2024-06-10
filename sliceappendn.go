package bootcamp

func SliceAppendN(n int) []int {
	slc := make([]int, 0)
	for i := 0; i < n; i++ {
		slc = append(slc, i)
	}

	return slc
}
