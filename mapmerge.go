package bootcamp

func MapMerge(m1 map[string]int, m2 map[string]int) map[string]int {
	m1_clone := MapClone1(m1)

	for key, val := range m2 {
		m1_clone[key] = val
	}

	return m1_clone
}

func MapClone1(m map[string]int) map[string]int {
	map_clone := make(map[string]int)

	for key, value := range m {
		map_clone[key] = value
	}

	return map_clone
}
