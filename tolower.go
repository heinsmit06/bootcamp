package bootcamp

func ToLower(s string) string {
	for i := 0; i < len(s); i++ {
		if 65 <= s[i] && s[i] <= 90 {
			s = s[:i] + string((s[i] + 32)) + s[(i+1):]
		}
	}

	return s
}
