package bootcamp

func UniqComb2(characters string) []string {
	if len(characters) <= 1 {
		return []string{}
	}

	unique := make(map[string]int)
	result := []string{}

	for i := range characters {
		if unique[string(characters[i])] > 0 {
			return []string{}
		} else {
			unique[string(characters[i])]++
		}
	}

	for i := 0; i < len(characters); i++ {
		for j := 0; j < len(characters); j++ {
			if characters[i] != characters[j] {
				result = append(result, string([]rune{rune(characters[i]), characters[j]}))
			}
		}
	}

	return result
}
