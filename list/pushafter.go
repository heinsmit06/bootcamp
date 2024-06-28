package list

func (l *List) PushAfter(n *ListNode, v interface{}) {
	temp := n.Next
	n.Next = &ListNode{Value: v, Next: temp}
}
