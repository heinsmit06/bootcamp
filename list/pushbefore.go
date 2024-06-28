package list

func (l *List) PushBefore(n *ListNode, v interface{}) {
	if l.Head == nil || n == nil {
		return
	}

	current := l.Head
	if current == n {
		temp := l.Head
		l.Head = &ListNode{Value: v, Next: temp}
		return
	}

	for current.Next != n && current.Next != nil {
		current = current.Next
	}

	if current.Next == n {
		l.PushAfter(current, v)
	} else if current.Next == nil {
		return
	}
}
