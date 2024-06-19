package bootcamp

func Split2(s string, sep string) []string {
	if sep == "" {
		// Special case: split into individual characters
		result := make([]string, len(s))
		for i := 0; i < len(s); i++ {
			result[i] = string(s[i])
		}
		return result
	}

	var result []string
	sepLen := len(sep)
	start := 0
	for i := 0; i < len(s); i++ {
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			start = i + sepLen
			i += sepLen - 1
		}
	}
	// Add the remaining part of the string
	result = append(result, s[start:])
	return result
}

func MorseToText(s string) string {
	str := Split2(s, " ")
	morse := map[string]rune{
		".-":     'A',
		"-...":   'B',
		"-.-.":   'C',
		"-..":    'D',
		".":      'E',
		"..-.":   'F',
		"--.":    'G',
		"....":   'H',
		"..":     'I',
		".---":   'J',
		"-.-":    'K',
		".-..":   'L',
		"--":     'M',
		"-.":     'N',
		"---":    'O',
		".--.":   'P',
		"--.-":   'Q',
		".-.":    'R',
		"...":    'S',
		"-":      'T',
		"..-":    'U',
		"...-":   'V',
		".--":    'W',
		"-..-":   'X',
		"-.--":   'Y',
		"--..":   'Z',
		".----":  '1',
		"..---":  '2',
		"...--":  '3',
		"....-":  '4',
		".....":  '5',
		"-....":  '6',
		"--...":  '7',
		"---..":  '8',
		"----.":  '9',
		"-----":  '0',
		".-.-.-": '.',
		"--..--": ',',
		"..--..": '?',
	}
	var res string
	for _, v := range str {
		res += string(morse[v])
	}
	return res
}
