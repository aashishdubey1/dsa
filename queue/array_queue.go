package queue

type ArrayQueue[T any] struct {
	items []T
}

func (q *ArrayQueue[T]) Enqueue(val T) {
	q.items = append(q.items, val)
}

func (q *ArrayQueue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var el T
		return el, false
	}
	el := q.items[0]
	q.items = q.items[1:]
	return el, true
}

func (q *ArrayQueue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *ArrayQueue[T]) Peek() (T, bool) {
	if len(q.items) == 0 {
		var res T
		return res, false
	}
	return q.items[0], true
}
