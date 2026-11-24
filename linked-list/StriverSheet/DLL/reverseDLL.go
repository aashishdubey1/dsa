package dll

func ReverseDLL(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	curr := head
	var newHead *ListNode

	for curr != nil {
		temp := curr.Next
		curr.Next = curr.Prev
		curr.Prev = temp

		newHead = curr
		curr = curr.Prev
	}
	return newHead
}
