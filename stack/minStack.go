package stack

type MinStack struct {
	mainStack ArrayStack
	minStack  ArrayStack
}

func (ms *MinStack) Push(val int) {
	if ms.minStack.IsEmpty() {
		ms.minStack.Push(val)
	} else {
		curr, _ := ms.minStack.Peek()
		if val <= curr {
			ms.minStack.Push(val)
		}
	}
	ms.mainStack.Push(val)
}

func (ms *MinStack) Pop() (int, bool) {
	val, ok := ms.mainStack.Pop()
	if !ok {
		return 0, false
	}

	curr, _ := ms.minStack.Peek()
	if val == curr {
		ms.minStack.Pop()
	}
	return val, true
}

func (ms *MinStack) GetMin() (int, bool) {
	return ms.minStack.Peek()
}
