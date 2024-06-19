package bootcamp

func MapEqual(m1 map[string]int, m2 map[string]int) bool {
	if len(m1) == len(m2) {
		for key, value := range m1 {
			if val, ok := m2[key]; !ok || val != value {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
