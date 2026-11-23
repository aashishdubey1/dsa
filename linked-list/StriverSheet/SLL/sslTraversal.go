package sll

// Given the head of a singly Linked List.
// Traverse the entire Linked List and return
// its elements in an array in the order of their appearance.

func SLLTraversal(list *SLL) []int {
	curr := list.Head
	result := []int{}

	for curr != nil {
		result = append(result, curr.Value)
		curr = curr.Next
	}
	return result
}
