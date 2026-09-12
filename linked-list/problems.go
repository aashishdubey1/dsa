package linkedlist

import ()

func ReverseList(list *LinkedList) {

	prev := (*Node)(nil)
	curr := list.Head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	list.Head = prev
}

func RunProblems() {

	list := LinkedList{}
	list.InsertAtHead(5)
	list.InsertAtHead(29)
	list.InsertAtHead(20)
	list.InsertAtTail(30)

	list.Print()

	ReverseList(&list)

	list.Print()
}
