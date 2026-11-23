package sll

func InsertAtKthPosition(list *SLL, k int, x int) *SLL {
	if list.Head == nil || k <= 0 {
		return list
	}

	newNode := &Node{Value: x}

	if k == 1 {
		newNode.Next = list.Head
		list.Head = newNode
		return list
	}

	curr := list.Head
	count := 1

	for count < k-1 && curr.Next != nil {
		curr = curr.Next
		count++
	}

	if count != k-1 {
		return list
	}

	newNode.Next = curr.Next
	curr.Next = newNode

	return list
}
