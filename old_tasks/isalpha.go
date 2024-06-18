package bootcamp

func IsAlpha(s string) bool {
	counter := 0
	for i := 0; i < len(s); i++ {
		if (65 <= s[i] && s[i] <= 90) || (97 <= s[i] && s[i] <= 122) {
			counter++
		}
	}

	if counter == len(s) {
		return true
	} else {
		return false
	}
}
