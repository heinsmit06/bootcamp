package bootcamp

func RomanToInt(roman string) int {
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// Helper function to check if a given Roman numeral is valid
	isValidRoman := func(s string) bool {
		validNumerals := "IVXLCDM"
		for i := 0; i < len(s); i++ {
			found := false
			for j := 0; j < len(validNumerals); j++ {
				if s[i] == validNumerals[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
		return true
	}

	if !isValidRoman(roman) {
		return 0
	}

	result := 0
	length := len(roman)
	for i := 0; i < length; i++ {
		value := romanValues[roman[i]]
		if i+1 < length && value < romanValues[roman[i+1]] {
			result -= value
		} else {
			result += value
		}
	}

	return result
}
