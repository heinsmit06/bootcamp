package bootcamp

func MapClone(m map[string]int) map[string]int {
	map_clone := make(map[string]int)

	for key, value := range m {
		map_clone[key] = value
	}

	return map_clone
}
