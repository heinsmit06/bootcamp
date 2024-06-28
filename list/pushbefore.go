package list

func (l *List) PushBefore(n *ListNode, v interface{}) {
	current := l.Head
	for current.Next != n && current.Next != nil {
		current = current.Next
	}

	l.PushAfter(current, v)
}
