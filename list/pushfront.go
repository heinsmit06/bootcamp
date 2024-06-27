package list

func (l *List) PushFront(v interface{}) {
	newNode := &ListNode{Value: v}
	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	l.Head = newNode
	newNode.Next = current
}
