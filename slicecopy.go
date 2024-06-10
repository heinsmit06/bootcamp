package bootcamp

func SliceCopy(dst, src []int) []int {
	for i := 0; i < len(dst); i++ {
		dst[i] = src[i]
	}

	return dst
}
