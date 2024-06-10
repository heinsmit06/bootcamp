package bootcamp

func SliceMakeN(n int) []int {
	slc := make([]int, n)
	for i := 0; i < len(slc); i++ {
		slc[i] = i
	}

	return slc
}
