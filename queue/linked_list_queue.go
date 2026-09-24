package queue

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Queue[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (q *Queue[T]) Enqueue(val T) {
	newNode := &Node[T]{Value: val}
	if q.Tail == nil {
		q.Head = newNode
		q.Tail = newNode
		return
	}
	q.Tail.Next = newNode
	q.Tail = newNode
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.Head == nil {
		var res T
		return res, false
	}
	node := q.Head
	q.Head = q.Head.Next
	if q.Head == nil {
		q.Tail = nil
	}
	return node.Value, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.Head == nil {
		var res T
		return res, false
	}
	return q.Head.Value, true
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Head == nil
}
