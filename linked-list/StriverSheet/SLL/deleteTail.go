package sll

func DeleteTail(list *SLL) *SLL {

	if list.Head == nil {
		return list
	}

	if list.Head.Next == nil {
		list.Head = nil
		return list
	}

	curr := list.Head

	for curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next = nil

	return list
}
