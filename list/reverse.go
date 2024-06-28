package list

func (l *List) Reverse() {
	if l.Head == nil {
		return
	}

	current := l.Head
	var previous *ListNode = nil

	for current != nil {
		current.Next = previous
		previous = current
		current = current.Next
	}

	l.Head = previous
}
