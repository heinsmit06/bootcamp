package bootcamp

func itoa(n int) string {
	neg := false
	res := []rune{}

	if n == 0 {
		return "0"
	} else if n < 0 {
		n *= -1
		neg = true
	}

	for n > 0 {
		digit := n % 10
		res = append([]rune{rune(digit + '0')}, res...)
		n /= 10
	}

	if neg {
		res = append([]rune{'-'}, res...)
	}

	return string(res)
}

func MagicGrowth2() []string {
	result := []string{}
	for i := 0; i < 9; i++ {
		for j := i + 1; j < 9; j++ {
			result = append(result, itoa(i)+itoa(j))
		}
	}

	return result
}
