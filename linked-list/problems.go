package linkedlist

func ReversList(linkedlist *LinkedList) {
	var prev *Node
	curr := linkedlist.Head

	for curr.Next != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	linkedlist.Head = prev

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

		if slow == fast {
			return true
		}
	}
	return false
}

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
