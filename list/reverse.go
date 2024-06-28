package list

func (l *List) Reverse() {
	if l.Head == nil {
		return
	}

	current := l.Head
	previous := l.Head
	for current != nil {
		if current == l.Head {
			continue
		}

		current.Next = previous

		previous = current
		current = current.Next
	}

	l.Head = current
}
