package dll

func InsertBeforeGivenNode(node *ListNode, x int) {

	newNode := &ListNode{Value: x}

	prev := node.Prev
	prev.Next = newNode
	newNode.Prev = prev

	newNode.Next = node
	node.Prev = newNode
}
