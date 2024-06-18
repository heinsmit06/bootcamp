package bootcamp

func SliceUnion(slices ...[]int) []int {
	slc := []int{}

	for i := 0; i < len(slices); i++ {
		for _, v := range slices[i] {
			slc = append(slc, v)
		}
	}

	m := make(map[int]bool)
	unique_slice := []int{}
	for _, val := range slc {
		if _, ok := m[val]; !ok {
			m[val] = true
			unique_slice = append(unique_slice, val)
		}
	}

	return unique_slice
}
