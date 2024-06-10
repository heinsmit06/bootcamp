package bootcamp

func SliceCopy(dst, src []int) []int {
	if len(dst) == 0 {
		return dst
	}

	for i := 0; i < len(dst); i++ {
		dst[i] = src[i]
	}

	return dst
}
