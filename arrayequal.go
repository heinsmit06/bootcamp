package bootcamp

func ArrayEqual(arr1, arr2 *[20]int) bool {
	counter := 0
	for i := 0; i < len(*arr1); i++ {
		if arr1[i] == arr2[i] {
			counter++
		}
	}

	if counter == 20 {
		return true
	} else {
		return false
	}
}
