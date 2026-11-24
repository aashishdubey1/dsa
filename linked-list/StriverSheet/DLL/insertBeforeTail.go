package dll

func InsertBeforeTail(head *ListNode, x int) *ListNode {

	newNode := &ListNode{Value: x}

	if head == nil {
		return newNode
	}

	curr := head

	for curr.Next != nil {
		curr = curr.Next
	}

	if curr.Prev != nil {
		curr.Prev.Next = newNode
		newNode.Prev = curr.Prev
	}
	newNode.Next = curr
	curr.Prev = newNode

	return head
}
