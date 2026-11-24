package dll

func ArrayToDLL(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Value: arr[0]}
	curr := head

	for i := 1; i < len(arr); i++ {
		newNode := &ListNode{Value: arr[i]}

		curr.Next = newNode
		newNode.Prev = curr

		curr = newNode
	}
	return head
}
