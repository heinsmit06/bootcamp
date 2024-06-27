package list

type ListNode struct {
	Next  *ListNode
	Value interface{}
}

type List struct {
	Head, Tail *ListNode
}

func NewList() *List {
	newList := &List{}
	listNodeHead, listNodeTail := &ListNode{}, &ListNode{}
	newList.Head = listNodeHead
	newList.Tail = listNodeTail
	return newList
}
