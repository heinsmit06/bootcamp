package bootcamp

func Replace(s string, old string, new string) string {
	if Contains(s, old) {
		return Join(Split(s, old), new)
	}

	return s
}
