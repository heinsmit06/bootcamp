package bootcamp

func Trim(s string) string {
	return ReverseString(RemoveLeadingSpaces(ReverseString(RemoveLeadingSpaces(s))))
}

func RemoveLeadingSpaces(s string) string {
	var res string
	is_space := 1

	for _, v := range s {
		if v != 32 || is_space == 0 {
			res += string(v)
			is_space = 0
		}
	}

	return res
}

func ReverseString(s string) string {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += string(rune(s[i]))
	}

	return res
}
