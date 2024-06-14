package bootcamp

func RoundBrackets(s string) bool {
	if len(s) == 0 {
		return false
	}

	open := 0
	for i := 0; i < len(s); i++ {
		if rune(s[i]) == ')' && open == 0 {
			return false
		} else if rune(s[i]) == '(' {
			open++
		} else if rune(s[i]) == ')' && open > 0 {
			open--
		}
	}

	if open > 0 {
		return false
	}

	return true
}
