package bootcamp

func IsNumeric(s string) bool {
	counter := 0
	if s[0] == '-' || s[0] == '+' {
		for i := 1; i < len(s); i++ {
			if 48 <= s[i] && s[i] <= 57 {
				counter++
			}
		}
	} else {
		for i := 0; i < len(s); i++ {
			if 48 <= s[i] && s[i] <= 57 {
				counter++
			}
		}
	}

	if (s[0] == '-' || s[0] == '+') && (counter == len(s)-1) {
		return true
	} else if counter == len(s) {
		return true
	} else {
		return false
	}
}
