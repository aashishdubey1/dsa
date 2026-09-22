package stack

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	Top  *Node
	size int
}

func (s *Stack) Push(val int) {
	newNode := &Node{Value: val}
	newNode.Next = s.Top
	s.Top = newNode
	s.size++
}
func (s *Stack) Pop() (int, bool) {
	if s.Top == nil {
		return -1, false
	}
	curr := s.Top
	s.Top = s.Top.Next
	s.size--
	return curr.Value, true
}

func (s *Stack) Peek() (int, bool) {
	if s.Top == nil {
		return 0, false
	}
	return s.Top.Value, true
}

func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}

//
// func (s *Stack) Size() int {
// 	count := 0
// 	for curr := s.Top; curr != nil; curr = curr.Next {
// 		count++
// 	}
// 	return count
// }

func (s *Stack) Size() int {
	return s.size
}
