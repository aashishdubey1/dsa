package stack

type ArrayStack struct {
	Items []int
}

func (as *ArrayStack) Push(val int) {
	as.Items = append(as.Items, val)
}

func (as *ArrayStack) Pop() (int, bool) {
	if len(as.Items) == 0 {
		return 0, false
	}
	index := len(as.Items) - 1
	val := as.Items[index]
	as.Items = as.Items[:index]
	return val, true
}

func (as *ArrayStack) Peek() (int, bool) {
	if len(as.Items) == 0 {
		return 0, false
	}
	return as.Items[len(as.Items)-1], true
}

func (as *ArrayStack) IsEmpty() bool {
	return len(as.Items) == 0
}

func (as *ArrayStack) Size() int {
	return len(as.Items)
}
