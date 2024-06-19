package bootcamp

func MapKeys(m map[string]int) []string {
	string_slice := make([]string, 0)

	for key := range m {
		string_slice = append(string_slice, key)
	}

	return string_slice
}
