package sll

func GetLength(list *SLL) int {
	curr := list.Head
	length := 0
	for curr != nil {
		curr = curr.Next
		length++
	}
	return length
}
