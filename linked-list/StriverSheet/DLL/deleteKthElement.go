package dll

func DeleteKthElement(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return head
	}

	curr := head
	count := 1

	for count < k && curr != nil {
		count++
		curr = curr.Next
	}

	if curr == nil {
		return head
	}

	if curr.Prev != nil {
		curr.Prev.Next = curr.Next
	} else {
		head = curr.Next
	}

	if curr.Next != nil {
		curr.Next.Prev = curr.Prev
	}

	curr.Prev = nil
	curr.Next = nil

	return head
}
