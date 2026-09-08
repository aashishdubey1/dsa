package trees

import "fmt"

type Node struct { 
	Value int 
	Left *Node 
	Right *Node 
}	


func PreOrder(node *Node) { 
	if node == nil { 
		return 
	}
	fmt.Println(node.Value)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

func PostOrder(node *Node) { 
	if node == nil {
		return 
	}
	PostOrder(node.Left)
	PostOrder(node.Right)
	fmt.Println(node.Value)
}

func InOrder(node *Node) { 
	if node == nil {
		return 
	}
	InOrder(node.Left)
	fmt.Println(node.Value)
	InOrder(node.Right)
}



func Run(){

	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	PreOrder(root)

}