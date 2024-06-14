package bootcamp

func IsPalindromeWord(s string) bool {
	temp := []rune(s)
	if len(temp) <= 0 {
		return false
	}

	for low, high := 0, len(s)-1; low < high; low, high = low+1, high-1 {
		if temp[high] >= 'A' && temp[high] <= 'Z' {
			temp[high] = rune(temp[high] + 32)
		}
		if temp[low] >= 'A' && temp[low] <= 'Z' {
			temp[low] = rune(temp[low] + 32)
		}
		if temp[low] < 'a' || temp[low] > 'z' {
			return false
		} else if temp[high] < 'a' || temp[high] > 'z' {
			return false
		} else if temp[low] != temp[high] {
			return false
		}
	}

	return true
}
