package linkedlist

func ReversList(list *LinkedList) {
	var prev *Node
	curr := list.Head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	list.Head = prev
}

func MiddleNode(list *LinkedList) *Node {
	slow := list.Head
	fast := list.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func HasCycle(list *LinkedList) bool {
	slow := list.Head
	fast := list.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// func MergeList(list1 *LinkedList, list2 *LinkedList) *LinkedList {
//
// 	c1 := list1.Head
// 	c2 := list2.Head
//
// 	for c1.Next != nil {
// 		if c1.Value < c2.Value {
// 			list2.Head = c2.Next
// 			c2.Next = c1.Next
// 			c1.Next = c2
// 		} else if c2.Value < c2.Value {
//
// 		}
//
// 	}
//
// }

func RunProblems() {

	list := LinkedList{}
	list.InsertAtHead(5)
	list.InsertAtHead(6)
	list.InsertAtHead(7)
	list.InsertAtHead(8)

	list.Print()

	// ReversList(&list)
	// list.Print()

	MiddleNode(&list)

}
