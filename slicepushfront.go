package bootcamp

func SlicePushFront(arr *[]int, values ...int) {
	slc := []int{}
	for _, value := range values {
		slc = append(slc, value)
	}

	for i := 0; i < len(*arr); i++ {
		slc = append(slc, (*arr)[i])
	}

	*arr = slc
}
