package dll

func InsertBeforeHead(head *ListNode, x int) *ListNode {
	newNode := &ListNode{Value: x}
	if head == nil {
		return newNode
	}

	newNode.Next = head
	head.Prev = newNode

	return newNode
}
