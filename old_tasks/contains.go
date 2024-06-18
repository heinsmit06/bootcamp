package bootcamp

func Contains(str string, substr string) bool {
	for i := 0; i < len(str); i++ {
		if (str[i] == substr[0]) && (i+len(substr) <= len(str)) {
			if str[i:i+len(substr)] == substr {
				return true
			}
		}
	}

	return false
}
