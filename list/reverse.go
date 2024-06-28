package list

func (l *List) Reverse() {
	if l.Head == nil {
		return
	}

	current := l.Head
	var previous *ListNode = nil
	var next *ListNode = nil

	for current != nil {
		next = current.Next
		current.Next = previous
		current.Prev = next
		previous = current
		current = next
	}

	l.Head = previous
}
