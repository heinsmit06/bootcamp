package list

func (l *List) Find(fn func(v interface{}) bool) *ListNode {
	if l.Head == nil {
		return nil
	}

	current := l.Head
	for current.Next != nil {
		if fn(current.Value) {
			return current
		}

		current = current.Next
	}

	return nil
}
