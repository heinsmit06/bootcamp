package list

func (l *List) Remove(n *ListNode) {
	if l.Head == nil || n == nil {
		return
	}

	if l.Head == n {
		l.Head = n.Next
		return
	}

	current := l.Head

	for current.Next != n && current.Next != nil {
		current = current.Next
	}

	current.Next = n.Next
}
