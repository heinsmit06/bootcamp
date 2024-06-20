package bootcamp

func MissingNumber(arr []int) int {
	n := 1
	flag := 0
	for n > 0 {
		for _, v := range arr {
			if n == v {
				flag = 1
			}
		}

		if flag == 1 {
			flag = 0
			n++
			continue
		} else {
			return n
		}
	}

	return 1
}
