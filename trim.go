package bootcamp

func Trim(s string) string {
	return ReverseString(RemoveLeadingSpaces(ReverseString(RemoveLeadingSpaces(s))))
}

func RemoveLeadingSpaces(s string) string {
	var res string

	for i, v := range s {
		if v != 32 {
			res += s[i:]
			break
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
