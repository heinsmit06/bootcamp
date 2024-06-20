package bootcamp

func ConvertBaseNbr(n string, base string) int {
	if len(base) < 2 || hasDuplicates(base) {
		return -1
	}

	baseMap := make(map[byte]int)
	for i := 0; i < len(base); i++ {
		baseMap[base[i]] = i
	}

	result := 0
	for i := 0; i < len(n); i++ {
		val, ok := baseMap[n[i]]
		if !ok {
			return -1
		}
		result = result*len(base) + val
	}

	return result
}

func hasDuplicates(s string) bool {
	seen := make(map[byte]struct{})
	for i := 0; i < len(s); i++ {
		if _, ok := seen[s[i]]; ok {
			return true
		}
		seen[s[i]] = struct{}{}
	}
	return false
}
