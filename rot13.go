package bootcamp

func Rot13(s string) string {
	var res string

	for _, v := range s {
		if (v >= 65 && v <= 77) || (v >= 97 && v <= 109) {
			res += string(v + 13)
		} else if (v >= 78 && v <= 90) || (v >= 110 && v <= 122) {
			res += string(v - 13)
		} else {
			res += string(v)
		}
	}
	return res
}
