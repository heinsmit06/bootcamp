package bootcamp

func Atoi(s string) int {
	var res int
	power := 0
	for i := len(s) - 1; i >= 0; i-- {
		res += int(s[i]-48) * PowRecursion(10, power)
		power++
	}

	return res
}
