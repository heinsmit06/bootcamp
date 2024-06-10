package bootcamp

func SliceCopy(dst, src []int) []int {
	if len(src) < len(dst) {
		for i := 0; i < len(src); i++ {
			dst[i] = src[i]
		}

		return dst
	}

	for j := 0; j < len(dst); j++ {
		dst[j] = src[j]
	}

	return dst
}
