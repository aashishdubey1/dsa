package sll

func Reverse(list *SLL) *SLL {
	var prev *Node
	curr := list.Head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	list.Head = prev
	return list
}

// Recursive
func ReverseRecursive(list *SLL) *SLL {
	if list.Head == nil {
		return list
	}

	list.Head = rev(list.Head)
	return list
}

func rev(curr *Node) *Node {
	if curr.Next == nil {
		return curr
	}
	node := rev(curr.Next)
	curr.Next.Next = curr
	curr.Next = nil
	return node
}
