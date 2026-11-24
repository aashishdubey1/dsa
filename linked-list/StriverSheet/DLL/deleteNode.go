package dll

func DeleteNode(node *ListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
