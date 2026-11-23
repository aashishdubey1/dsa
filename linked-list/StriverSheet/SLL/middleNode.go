package sll

//
// func MiddleNode(list *SLL) *Node {
// 	length := 0
// 	curr := list.Head
//
// 	for curr != nil {
// 		curr = curr.Next
// 		length++
// 	}
//
// 	curr = list.Head
// 	for range length / 2 {
// 		curr = curr.Next
// 	}
// 	return curr
// }

// Slow and Fast pointer approach

func MiddleNode(list *SLL) *Node {

	slow := list.Head
	fast := list.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
