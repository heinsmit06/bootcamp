package bootcamp

func MapLen(m map[string]int) int {
	length := 0
	for range m {
		length++
	}

	return length
}
