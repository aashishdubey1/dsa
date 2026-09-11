package linkedlist

import "fmt"

// func PrintList(node *Node) {
// 	curr := node
// 	for curr != nil {
// 		fmt.Println(curr.Value)
// 		curr = curr.Next
// 	}
// }

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Print() {
	curr := list.Head
	for curr != nil {
		fmt.Println(curr.Value)
		curr = curr.Next
	}
}

func (list *LinkedList) InsertAtHead(value int) {
	newNode := &Node{Value: value}
	newNode.Next = list.Head
	list.Head = newNode
}

func (list *LinkedList) InsertAtTail(value int) {
	newNode := &Node{Value: value}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	curr := list.Head

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNode
}

func RunList() {

	list := LinkedList{}

	a := &Node{Value: 6}
	b := &Node{Value: 8}
	c := &Node{Value: 12}

	a.Next = b
	b.Next = c
	c.Next = nil

	list.Head = a

	// list.Print()
	// list.InsertAtHead(5)
	// list.InsertAtHead(10)
	// list.InsertAtHead(19)
	list.InsertAtHead(2)
	list.InsertAtTail(5)
	list.Print()

}
