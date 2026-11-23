package sll

func InsertAtTail(list *SLL, x int) *SLL {

	newNode := &Node{Value: x}

	if list.Head == nil {
		list.Head = newNode
		return list
	}

	curr := list.Head

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNode

	return list
}
