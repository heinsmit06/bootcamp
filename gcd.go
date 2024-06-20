package bootcamp

func GCD(a, b int) int {
	for i := a; i >= 0; i-- {
		if (a%i == 0) && (b%i == 0) {
			return i
		}
	}

	return 0
}
