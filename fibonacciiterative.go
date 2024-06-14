package bootcamp

func FibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}
	n2, n1 := 0, 1
	for i := 2; i <= n; i++ {
		n2, n1 = n1, n1+n2
	}
	return n1
}
