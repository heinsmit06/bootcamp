package bootcamp

func MapValues(m map[string]int) []int {
	var int_slice []int

	for _, val := range m {
		int_slice = append(int_slice, val)
	}

	return int_slice
}
