package sll

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type SLL struct {
	Head *Node
}

func RunSSL() {
	list := SLL{}
	// list.Head.Next = &Node{Value: 20}
	// list.Head.Next.Next = &Node{Value: 30}
	// list.Head.Next.Next.Next = &Node{Value: 40}
	// list.Head.Next.Next.Next.Next = &Node{Value: 50}
	InsertAtHead(&list, 50)
	InsertAtHead(&list, 40)
	InsertAtHead(&list, 30)
	InsertAtHead(&list, 20)
	InsertAtHead(&list, 10)

	fmt.Println(SLLTraversal(&list))
}
