package sll

func DeletetKthNode(list *SLL, k int) *SLL {
	if list.Head == nil {
		return list
	}
	if k == 1 {
		list.Head = list.Head.Next
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

	dn := curr.Next
	curr.Next = dn.Next

	return list
}
