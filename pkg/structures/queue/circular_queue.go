package queue

import "fmt"

// CircularQueue implements a fixed-size circular queue.
type CircularQueue[T any] struct {
	items []T
	head  int
	tail  int
	size  int
	cap   int
}

// NewCircularQueue creates a new circular queue with the given capacity.
func NewCircularQueue[T any](capacity int) *CircularQueue[T] {
	return &CircularQueue[T]{
		items: make([]T, capacity),
		head:  0,
		tail:  0,
		size:  0,
		cap:   capacity,
	}
}

func (q *CircularQueue[T]) Enqueue(value T) error {
	if q.IsFull() {
		return fmt.Errorf("queue is full")
	}
	q.items[q.tail] = value
	q.tail = (q.tail + 1) % q.cap
	q.size++
	return nil
}

func (q *CircularQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	val := q.items[q.head]
	q.head = (q.head + 1) % q.cap
	q.size--
	return val, nil
}

func (q *CircularQueue[T]) IsFull() bool {
	return q.size == q.cap
}

func (q *CircularQueue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *CircularQueue[T]) Size() int {
	return q.size
}