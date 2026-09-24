package queue

type CircularBufferQueue[T any] struct {
	items    []T
	head     int
	tail     int
	capacity int
	size     int
}

func (q *CircularBufferQueue[T]) Enqueue(val T) bool {
	if q.capacity == q.size {
		return false
	}
	q.items[q.tail] = val
	q.tail = (q.tail + 1) % q.capacity
	q.size++
	return true
}

func (q *CircularBufferQueue[T]) Dequeue() (T, bool) {
	var zero T
	if q.size == 0 {
		return zero, false
	}
	val := q.items[q.head]
	q.items[q.head] = zero
	q.head = (q.head + 1) % q.capacity
	q.size--
	return val, true
}
func NewCircularBufferQueue[T any](capacity int) *CircularBufferQueue[T] {
	return &CircularBufferQueue[T]{
		items:    make([]T, capacity),
		capacity: capacity,
	}
}
