package list

import "fmt"

// Node represents a node in the linked list.
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// LinkedList represents a singly linked list.
type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// New creates a new empty LinkedList.
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// PushFront adds a value to the front of the list.
func (l *LinkedList[T]) PushFront(value T) {
	newNode := &Node[T]{Value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.Next = l.head
		l.head = newNode
	}
	l.size++
}

// PushBack adds a value to the back of the list.
func (l *LinkedList[T]) PushBack(value T) {
	newNode := &Node[T]{Value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.Next = newNode
		l.tail = newNode
	}
	l.size++
}

// Size returns the number of elements in the list.
func (l *LinkedList[T]) Size() int {
	return l.size
}

// Get returns the value at the specified index.
func (l *LinkedList[T]) Get(index int) (T, error) {
	if index < 0 || index >= l.size {
		var zero T
		return zero, fmt.Errorf("index out of bounds")
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value, nil
}

// InsertAt inserts a value at the specified index.
func (l *LinkedList[T]) InsertAt(index int, value T) error {
	if index < 0 || index > l.size {
		return fmt.Errorf("index out of bounds")
	}

	if index == 0 {
		l.PushFront(value)
		return nil
	}

	if index == l.size {
		l.PushBack(value)
		return nil
	}

	newNode := &Node[T]{Value: value}
	prev := l.head
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}

	newNode.Next = prev.Next
	prev.Next = newNode
	l.size++
	return nil
}

// Remove removes the element at the specified index and returns its value.
func (l *LinkedList[T]) Remove(index int) (T, error) {
	if index < 0 || index >= l.size {
		var zero T
		return zero, fmt.Errorf("index out of bounds")
	}

	var removeVal T

	if index == 0 {
		removeNode := l.head
		l.head = l.head.Next
		if l.head == nil {
			l.tail = nil
		}
		l.size--
		return removeNode.Value, nil
	}

	prev := l.head
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}

	removeNode := prev.Next
	prev.Next = removeNode.Next

	if index == l.size-1 {
		l.tail = prev
	}

	l.size--
	removeVal = removeNode.Value
	return removeVal, nil
}

// Clear removes all elements from the list.
func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}