package bootcamp

func Eraser(s string) string {
	slc := make([]rune, 0)

	for _, v := range s {
		if v != '<' {
			slc = append(slc, v)
		} else if len(slc) > 0 {
			slc = slc[:len(slc)-1]
		}
	}

	return string(slc)
}
