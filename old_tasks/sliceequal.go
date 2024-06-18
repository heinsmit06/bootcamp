package bootcamp

func SliceEqual(arr1, arr2 []int) bool {
	counter := 0
	if len(arr1) == len(arr2) {
		for i := 0; i < len(arr1); i++ {
			if arr1[i] == arr2[i] {
				counter++
			}
		}
	}

	if counter == len(arr1) {
		return true
	}

	return false
}
