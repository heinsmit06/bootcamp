package bootcamp

import "unicode/utf8"

func CountRunes(s string) int {
	length := utf8.RuneCountInString(s)
	return length
}
