package bootcamp

func SliceMakeN(n int) []int {
	slc := []int{}
	for i := 0; i < n; i++ {
		slc = append(slc, i)
	}

	return slc
}
