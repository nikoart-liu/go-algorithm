package stack

import (
	"fmt"
	"go-algorithm/pkg/structures/list"
)

// Stack defines the interface for a stack.
type Stack[T any] interface {
	Push(value T)
	Pop() (T, error)
	Peek() (T, error)
	IsEmpty() bool
	Size() int
}

// SliceStack implements a stack using a slice.
type SliceStack[T any] struct {
	items []T
}

func NewSliceStack[T any]() *SliceStack[T] {
	return &SliceStack[T]{}
}

func (s *SliceStack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *SliceStack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	index := len(s.items) - 1
	val := s.items[index]
	s.items = s.items[:index]
	return val, nil
}

func (s *SliceStack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *SliceStack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *SliceStack[T]) Size() int {
	return len(s.items)
}

// ListStack implements a stack using a singly linked list.
type ListStack[T any] struct {
	list *list.LinkedList[T]
}

func NewListStack[T any]() *ListStack[T] {
	return &ListStack[T]{list: list.New[T]()}
}

func (s *ListStack[T]) Push(value T) {
	s.list.PushFront(value)
}

func (s *ListStack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	return s.list.Remove(0)
}

func (s *ListStack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	return s.list.Get(0)
}

func (s *ListStack[T]) IsEmpty() bool {
	return s.list.Size() == 0
}

func (s *ListStack[T]) Size() int {
	return s.list.Size()
}