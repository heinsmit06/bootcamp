package bootcamp

func TestNthRune(fn func(s string, n int) rune) bool {
	fl1 := false
	fl2 := false
	fl3 := false
	s1 := ""
	s2 := "abc"
	run := fn(s1, 8)
	if run == -1 {
		fl1 = true
	}
	run2 := fn(s2, 10)

	if run2 == -1 {
		fl2 = true
	}
	if fn(s2, 2) == 'b' {
		fl3 = true
	}
	return fl1 && fl2 && fl3
}
