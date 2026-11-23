package sll

func DeleteHead(list *SLL) *SLL {
	if list.Head == nil {
		return list
	}

	list.Head = list.Head.Next
	return list
}
