package list

func (l *List) Remove(n *ListNode) {
	current := l.Head
	for current.Next != n {
		current = current.Next
	}
	if n.Next != nil {
		current.Next = n.Next
	} else {
		current.Next = nil
	}
}
