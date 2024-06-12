package bootcamp

func LastRune(s string) rune {
	if len(s) == 0 {
		return rune(0)
	}

	return rune(s[len(s)-1])
}
