package bootcamp

func Sqrt(x int) int {
	if x < 0 {
		return -1
	}

	low, high := 0, x

	for low <= high {
		mid := low + (high-low)/2

		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
