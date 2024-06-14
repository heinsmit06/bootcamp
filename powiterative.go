package bootcamp

func PowIterative(x int, power int) int {
	if power < 0 {
		return -1
	} else if power == 0 {
		return x
	}

	result := 1
	for i := 0; i < power; i++ {
		result *= x
	}
	return result
}
