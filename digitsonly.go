package bootcamp

func DigitsOnly(s string) string {
	var new_string string

	for i := 0; i < len(s); i++ {
		if 48 <= s[i] && s[i] <= 57 {
			new_string = new_string + string(s[i])
		}
	}

	return new_string
}
