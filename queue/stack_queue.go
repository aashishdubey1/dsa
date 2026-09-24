package queue

type QueueFromStack[T any] struct {
	inStack  []T
	outStack []T
}

func (q *QueueFromStack[T]) Enqueue(val T) {
	q.inStack = append(q.inStack, val)
}

func (q *QueueFromStack[T]) Dequeue() (T, bool) {
	if len(q.outStack) == 0 && len(q.inStack) == 0 {
		var zero T
		return zero, false
	}
	if len(q.outStack) == 0 {
		for len(q.inStack) > 0 {
			el := q.inStack[len(q.inStack)-1]
			q.inStack = q.inStack[:len(q.inStack)-1]
			q.outStack = append(q.outStack, el)
		}
	}
	val := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return val, true
}
