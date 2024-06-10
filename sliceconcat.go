package bootcamp

func SliceConcat(slices ...[]int) []int {
	slc := make([]int, 0)
	for i := 0; i < len(slices); i++ {
		for _, num := range slices[i] {
			slc = append(slc, num)
		}
	}

	return slc
}
