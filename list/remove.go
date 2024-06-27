package list

func (l *List) Remove(n *ListNode) {
	if l.Head == n {
		l.Head = nil
		return
	}

	current := l.Head

	for current.Next != n && current.Next != nil {
		current = current.Next
	}

	if n.Next != nil {
		current.Next = n.Next
	} else {
		current.Next = nil
	}
}
