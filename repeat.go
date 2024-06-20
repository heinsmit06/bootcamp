package bootcamp

func Repeat(s string, count int) string {
	var result_str string

	for i := 0; i < count; i++ {
		result_str = result_str + s
	}

	return result_str
}
