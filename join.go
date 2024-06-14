package bootcamp

func Join(elements []string, sep string) string {
	var result_str string
	for i := 0; i < len(elements); i++ {
		result_str += elements[i]
		if i < len(elements)-1 {
			result_str += sep
		}
	}

	return result_str
}
