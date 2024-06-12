package bootcamp

func Capitalize(s string) string {
	for i := 0; i < len(s); i++ {
		if 65 <= s[i] && s[i] <= 90 {
			s = s[:i] + string((s[i] + 32)) + s[(i+1):]
		}
	}

	if 97 <= s[0] && s[0] <= 122 { // make it uppercase only if it is a lowercase letter character
		s = string((s[0] - 32)) + s[1:] // making the first letter uppercase
	}

	for j := 1; j < len(s); j++ {
		if s[j-1] == ' ' {
			s = s[:j] + string((s[j] - 32)) + s[(j+1):]
		}
	}

	return s

	//for i := 0; i < len(s); i++ {
	//	if (i == 0) && (97 <= s[i] && s[i] <= 122) {
	//		s = s[:i] + string((s[i] - 32)) + s[(i+1):]
	//		continue
	//	} else if (i == 0) && (65 <= s[i] && s[i] <= 90) {
	//		continue
	//	}
	//
	//	if 65 <= s[i] && s[i] <= 90 {
	//		s = s[:i] + string((s[i] + 32)) + s[(i+1):]
	//		continue
	//	}
	//
	//	if s[i] == ' ' {
	//		if (i + 1) != len(s) {
	//			if 97 <= s[i+1] && s[i+1] <= 122 {
	//				s = s[:i] + string((s[i] - 32)) + s[(i+1):]
	//			}
	//		}
	//		i++
	//	}
	//}
	//
	//return s
}
