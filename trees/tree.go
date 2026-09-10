package trees

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
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

func LevelOrder(node *Node) {
	if node == nil {
		return
	}
	queue := []*Node{node}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Println(current.Value)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

}

func ItretivePreOrder(node *Node) {
	if node == nil {
		return
	}
	stack := []*Node{node}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
		fmt.Println(curr.Value)
	}
}

func Run() {

	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	PreOrder(root)

}
