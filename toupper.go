package bootcamp

func ToUpper(s string) string {
	for i := 0; i < len(s); i++ {
		if 97 <= s[i] && s[i] <= 122 {
			s = s[:i] + string((s[i] - 32)) + s[(i+1):]
		}
	}

	return s
}
