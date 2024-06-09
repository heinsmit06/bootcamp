package bootcamp

func Swap(a, b *int) {
	if a == nil && b == nil {
		return
	} else if (a == nil && b != nil) || (a != nil && b == nil) {
		return
	}

	var temp int
	temp = *a
	*a = *b
	*b = temp
}
