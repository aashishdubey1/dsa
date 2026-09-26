package trees

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BST struct {
	Root *TreeNode
}

func insert(node *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{Value: val}
	if node == nil {
		return newNode
	}

	if val > node.Value {
		node.Right = insert(node.Right, val)
	} else if val < node.Value {
		node.Left = insert(node.Left, val)
	}
	return node
}

func (bst *BST) Insert(val int) {
	bst.Root = insert(bst.Root, val)
}

func search(node *TreeNode, val int) bool {
	if node == nil {
		return false
	}
	if node.Value == val {
		return true
	} else if node.Value < val {
		return search(node.Right, val)
	} else {
		return search(node.Left, val)
	}
}
func (bst *BST) Search(val int) bool {
	return search(bst.Root, val)
}
