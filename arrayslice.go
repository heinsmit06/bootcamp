package bootcamp

func ArraySlice(arr [20]int, low int, high int) []int {
	if (low < high) && (low >= 0 && high <= 20) {

		slice_length := high - low
		slc := make([]int, slice_length)
		slc_counter := 0

		for i := low; i < high; i++ {
			slc[slc_counter] = arr[i]
			slc_counter++
		}

		return slc

	} else {
		slc := []int{}
		return slc
	}
}
