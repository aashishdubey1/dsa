package dll

func DeleteTail(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return nil
	}

	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Prev.Next = nil
	curr.Prev = nil
	return head
}
