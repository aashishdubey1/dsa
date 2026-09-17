package doublylinkedlist

import "fmt"

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type DLL struct {
	Head *Node
}

func (list *DLL) Print() {
	curr := list.Head
	for curr != nil {
		fmt.Println(curr.Value)
		curr = curr.Next
	}
}
func (list *DLL) PrintReverse() {
	if list.Head == nil {
		return
	}
	curr := list.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	for curr != nil {
		fmt.Println(curr.Value)
		curr = curr.Prev
	}
}

func (list *DLL) InsertFirst(val int) {
	newNode := &Node{Value: val}
	if list.Head == nil {
		list.Head = newNode
		return
	}
	newNode.Next = list.Head
	list.Head.Prev = newNode
	list.Head = newNode
}
func (list *DLL) InsertLast(val int) {
	newNode := &Node{Value: val}
	if list.Head == nil {
		list.Head = newNode
		return
	}
	curr := list.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	newNode.Prev = curr
	curr.Next = newNode
}

func (list *DLL) InsertAt(index int, value int) {
	if index < 0 {
		fmt.Println("index out of bound")
		return
	}
	newNode := &Node{Value: value}
	if index == 0 {
		newNode.Next = list.Head
		if list.Head != nil {
			list.Head.Prev = newNode
		}
		list.Head = newNode
		return
	}

	curr := list.Head
	for range index - 1 {
		if curr == nil {
			fmt.Println("Index out of bound")
			return
		}
		curr = curr.Next
	}
	if curr == nil {
		fmt.Println("index out of bound ")
		return
	}
	newNode.Next = curr.Next
	if curr.Next != nil {
		curr.Next.Prev = newNode
	}
	curr.Next = newNode
	newNode.Prev = curr
}

func RunDLL() {

	list := DLL{}

	list.InsertFirst(5)
	list.InsertFirst(8)
	list.InsertFirst(9)
	list.InsertFirst(10)
	list.Print()
	list.InsertLast(25)
	list.Print()
}
