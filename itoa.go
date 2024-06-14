package bootcamp

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var result string
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	for n > 0 {
		digit := n % 10
		result = string(rune(digit+48)) + result
		n /= 10
	}

	return sign + result
}
