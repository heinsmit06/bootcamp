package list

func (l *List) Remove(n *ListNode) {
	if l.Head == nil || n == nil {
		return
	}

	if l.Head == n {
		l.Head = nil
		return
	}

	current := l.Head

	for current.Next != n && current.Next != nil {
		current = current.Next
	}
	if current.Next != nil {
		current.Next = n.Next
	}
}
