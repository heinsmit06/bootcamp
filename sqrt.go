package bootcamp

func Sqrt(x int) int {
	z := 1
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z
}
