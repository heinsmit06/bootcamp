package bootcamp

func CountRunes(s string) int {
	counter := 0
	for range s {
		counter++
	}

	return counter

	// length := utf8.RuneCountInString(s)
	// return length
}
