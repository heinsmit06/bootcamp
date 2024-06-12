package bootcamp

func Concat(s ...string) string {
	var result_str string

	for _, str := range s {
		result_str = result_str + str[:len(str)]
	}

	return result_str
}
