package bootcamp

func ToUpperRune(s rune) rune {
	if s >= 97 && s <= 122 {
		return s - 32
	}
	return s
}

func TextToMorse(s string) string {
	morse := map[rune]string{
		'A': ".-",
		'B': "-...",
		'C': "-.-.",
		'D': "-..",
		'E': ".",
		'F': "..-.",
		'G': "--.",
		'H': "....",
		'I': "..",
		'J': ".---",
		'K': "-.-",
		'L': ".-..",
		'M': "--",
		'N': "-.",
		'O': "---",
		'P': ".--.",
		'Q': "--.-",
		'R': ".-.",
		'S': "...",
		'T': "-",
		'U': "..-",
		'V': "...-",
		'W': ".--",
		'X': "-..-",
		'Y': "-.--",
		'Z': "--..",
		'1': ".----",
		'2': "..---",
		'3': "...--",
		'4': "....-",
		'5': ".....",
		'6': "-....",
		'7': "--...",
		'8': "---..",
		'9': "----.",
		'0': "-----",
		'.': ".-.-.-",
		',': "--..--",
		'?': "..--..",
	}
	var res string

	for i, v := range s {
		if v == ' ' {
			continue
		}
		if i == len(s)-1 {
			res += string(morse[ToUpperRune(v)])
			return res

		}
		res += string(morse[ToUpperRune(v)]) + " "

	}
	return res
}

// actual: "... .- .-.. . -- --..--  .... --- .--  .- .-. .  -.-- --- ..- ..--.."
// expected:"... .- .-.. . -- --..-- .... --- .-- .- .-. . -.-- --- ..- ..--.."
//          ... .- .-.. . -- --..-- .... --- .-- .- .-. . -.-- --- ..- ..--..
