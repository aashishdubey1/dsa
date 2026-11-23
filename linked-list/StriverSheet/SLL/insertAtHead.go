package sll

func InsertAtHead(list *SLL, x int) *SLL {
	newNode := &Node{Value: x}

	newNode.Next = list.Head
	list.Head = newNode

	return list
}
