package dll

func DeleteHead(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	newHead := head.Next

	if newHead != nil {
		newHead.Prev = nil
	}

	head.Next = nil
	return newHead
}
