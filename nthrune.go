package bootcamp

func NthRune(s string, n int) rune {
	if len(s) == 0 {
		return rune(0)
	}

	return rune(s[n])
}
