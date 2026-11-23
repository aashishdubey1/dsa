package sll

func DeleteNodeWithValue(list *SLL, val int) *SLL {

	if list.Head == nil {
		return list
	}

	if list.Head.Value == val {
		list.Head = list.Head.Next
		return list
	}

	curr := list.Head

	for curr.Next != nil && curr.Next.Value != val {
		curr = curr.Next
	}

	if curr.Next == nil {
		return list
	}

	curr.Next = curr.Next.Next

	return list
}
