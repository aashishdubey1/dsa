package dll

func InsertBeforeKthPosition(head *ListNode, x int, k int) *ListNode {

	if k <= 0 {
		return nil
	}

	newNode := &ListNode{Value: x}

	if head == nil {
		if k == 1 {
			return newNode
		}
		return head
	}

	if k == 1 {
		newNode.Next = head
		head.Prev = newNode
		return newNode
	}

	curr := head
	count := 1

	for count < k && curr != nil {
		curr = curr.Next
		count++
	}

	if curr == nil {
		return nil
	}

	curr.Prev.Next = newNode
	newNode.Prev = curr.Prev
	newNode.Next = curr
	curr.Prev = newNode
	return head
}
