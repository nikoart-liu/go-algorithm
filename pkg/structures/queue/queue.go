package queue

import (
	"fmt"
	"go-algorithm/pkg/structures/list"
)

// Queue defines the interface for a queue.
type Queue[T any] interface {
	Enqueue(value T)
	Dequeue() (T, error)
	Peek() (T, error)
	IsEmpty() bool
	Size() int
}

// SliceQueue implements a queue using a slice.
type SliceQueue[T any] struct {
	items []T
}

func NewSliceQueue[T any]() *SliceQueue[T] {
	return &SliceQueue[T]{}
}

func (q *SliceQueue[T]) Enqueue(value T) {
	q.items = append(q.items, value)
}

func (q *SliceQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	val := q.items[0]
	q.items = q.items[1:]
	// Optional: shrink underlying array if too much empty space
	return val, nil
}

func (q *SliceQueue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	return q.items[0], nil
}

func (q *SliceQueue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *SliceQueue[T]) Size() int {
	return len(q.items)
}

// ListQueue implements a queue using a singly linked list.
type ListQueue[T any] struct {
	list *list.LinkedList[T]
}

func NewListQueue[T any]() *ListQueue[T] {
	return &ListQueue[T]{list: list.New[T]()}
}

func (q *ListQueue[T]) Enqueue(value T) {
	q.list.PushBack(value)
}

func (q *ListQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	return q.list.Remove(0)
}

func (q *ListQueue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	return q.list.Get(0)
}

func (q *ListQueue[T]) IsEmpty() bool {
	return q.list.Size() == 0
}

func (q *ListQueue[T]) Size() int {
	return q.list.Size()
}