package list

// Sort sorts the nodes in the doubly linked list based on the comparison function provided.
func (l *List) Sort(fn func(a *ListNode, b *ListNode) int) {
	if l.Head == nil || l.Head.Next == nil {
		return
	}

	// Merge sort
	l.Head = mergeSort(l.Head, fn)

	// Fix the tail pointer
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	l.Tail = current
}

// mergeSort performs the merge sort on the linked list.
func mergeSort(head *ListNode, fn func(a *ListNode, b *ListNode) int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Split the list into two halves
	middle := getMiddle(head)
	nextOfMiddle := middle.Next
	middle.Next = nil

	// Recursively sort the two halves
	left := mergeSort(head, fn)
	right := mergeSort(nextOfMiddle, fn)

	// Merge the sorted halves
	return sortedMerge(left, right, fn)
}

// getMiddle finds the middle node of the linked list.
func getMiddle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// sortedMerge merges two sorted linked lists.
func sortedMerge(a *ListNode, b *ListNode, fn func(a *ListNode, b *ListNode) int) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	var result *ListNode

	if fn(a, b) <= 0 {
		result = a
		result.Next = sortedMerge(a.Next, b, fn)
	} else {
		result = b
		result.Next = sortedMerge(a, b.Next, fn)
	}

	return result
}
