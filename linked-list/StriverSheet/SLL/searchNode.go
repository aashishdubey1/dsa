package sll

func SearchNode(list *SLL, k int) bool {
	curr := list.Head
	for curr != nil {
		if curr.Value == k {
			return true
		}
		curr = curr.Next
	}
	return false
}
