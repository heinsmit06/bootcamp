package list

func (l *List) Len() int {
	if l.Head == nil {
		return 0
	}

	current := l.Head
	count := 0
	for current.Next != nil {
		count++
	}

	return count
}
