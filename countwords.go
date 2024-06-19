package bootcamp

func CountWords(s string) map[string]int {
	m := map[string]int{}
	lowercase_string := ToLower1(s)
	s_sliced := Split1(lowercase_string, " ")

	for i := 0; i < len(s_sliced); i++ {
		word := s_sliced[i]
		last_char := len(word) - 1
		if word[last_char] == '.' {
			s_sliced[i] = word[:last_char]
		}
	}

	for _, val := range s_sliced {
		m[val] += 1
	}

	return m
}

func ToLower1(s string) string {
	for i := 0; i < len(s); i++ {
		if 65 <= s[i] && s[i] <= 90 {
			s = s[:i] + string((s[i] + 32)) + s[(i+1):]
		}
	}

	return s
}

func Split1(s string, sep string) []string {
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
