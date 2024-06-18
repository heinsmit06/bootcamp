package bootcamp

func SliceSort(arr []int) {
	arr_length := len(arr)
	temp := 0

	for i := 0; i < arr_length-1; i++ {
		for j := 0; j < arr_length-i-1; j++ {
			if (arr)[j] > (arr)[j+1] {
				temp = (arr)[j]
				(arr)[j] = (arr)[j+1]
				(arr)[j+1] = temp
			}
		}
	}
}
