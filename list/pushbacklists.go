package list

func (l *List) PushBackLists(lists ...*List) {
	if l.Head == nil {
		return
	}

	for i := 0; i < len(lists); i++ {
		lastNode := l.Back()
		lastNode.Next = lists[i].Head
	}
}
