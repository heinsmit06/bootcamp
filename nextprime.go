package bootcamp

func IsPrime1(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func NextPrime(n int) int {
	if n <= 1 {
		return 2
	}

	i := n + 1
	for {
		if IsPrime1(i) {
			return i
		}

		i++
	}
}
