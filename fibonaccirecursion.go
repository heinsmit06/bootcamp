package bootcamp

func FibonacciRecursion(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 1
	}

	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}
