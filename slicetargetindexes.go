package bootcamp

func SliceTargetIndexes(arr []int, target int) []int {
	slc := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			slc = append(slc, i)
		}
	}
	return slc
}
