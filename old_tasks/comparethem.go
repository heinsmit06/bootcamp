package bootcamp

func CompareThem(a, b *int) bool {
	if a == nil && b == nil {
		return true
	} else if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}

	var a_value int = *a
	var b_value int = *b

	if a_value == b_value {
		return true
	} else {
		return false
	}
}
