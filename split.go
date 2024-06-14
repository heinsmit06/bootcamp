package bootcamp

func Split(s string, sep string) []string {
	str_slice := []string{}
	var start int = 0

	if sep == "" {
		for _, c := range s {
			str_slice = append(str_slice, string(c))
		}

		return str_slice
	} else {
		for i := 0; i < len(s); i++ {
			if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
				str_slice = append(str_slice, s[start:i])
				start = i + len(sep)
			}
		}
		str_slice = append(str_slice, s[start:])
	}

	return str_slice
}
