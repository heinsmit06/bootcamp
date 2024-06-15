package bootcamp

func combRecursive(chars []rune, n int, current string, combinations *[]string) {
	if n == 0 {
		*combinations = append(*combinations, current)
		return
	}

	for i, char := range chars {
		remaining := append([]rune{}, chars[:i]...)
		remaining = append(remaining, chars[i+1:]...)
		combRecursive(remaining, n-1, current+string(char), combinations)
	}
}

func UniqCombN(characters string, n int) []string {
	if n < 1 {
		return []string{}
	} else if len(characters) < n {
		return []string{}
	} else if len(characters) == 0 && n == 1 {
		return []string{""}
	}

	unique := make(map[rune]bool)
	for _, char := range characters {
		if unique[char] {
			return []string{}
		}
		unique[char] = true
	}

	var result []string
	combRecursive([]rune(characters), n, "", &result)

	return result
}
