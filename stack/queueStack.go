package stack

import "github.com/aashishdubey1/dsa/queue"

type StackFromQueue[T any] struct {
	q1 *queue.ArrayQueue[T]
	q2 *queue.ArrayQueue[T]
}

func (s *StackFromQueue[T]) Push(val T) {
	s.q2.Enqueue(val)
	if !s.q1.IsEmpty() {
		val, _ := s.q1.Dequeue()
		s.q2.Enqueue(val)
	}
	s.q1, s.q2 = s.q2, s.q1
}

func (s *StackFromQueue[T]) Pop() (T, bool) {
	if s.q1.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.q1.Dequeue()
}
