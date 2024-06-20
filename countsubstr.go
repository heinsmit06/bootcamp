package bootcamp

func CountSubstr(s string, substr string) int {
	var count int = Contains1(s, substr)
	return count
}

func Contains1(str string, substr string) int {
	count := 0
	for i := 0; i < len(str); i++ {
		if (str[i] == substr[0]) && (i+len(substr) <= len(str)) {
			if str[i:i+len(substr)] == substr {
				count++
			}
		}
	}

	return count
}
