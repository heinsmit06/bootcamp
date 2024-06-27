package list

func (l *List) Len() int {
	if l.Head == nil {
		return 0
	}

	current := l.Head
	count := 1
	for current.Next != nil {
		count++
		current = current.Next
	}

	return count
}
