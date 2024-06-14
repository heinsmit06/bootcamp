package bootcamp

var digits = "0123456789"

func helper(start int, path []rune, combinations *[]string, n int) {
	if len(path) == n {
		*combinations = append(*combinations, string(path))
		return
	}
	for i := start; i < len(digits); i++ {
		helper(i+1, append(path, rune(digits[i])), combinations, n)
	}
}

func MagicGrowthN(n int) []string {
	if n < 1 || n > 10 {
		return []string{}
	}
	combinations := []string{}
	helper(0, []rune{}, &combinations, n)

	return combinations
}
