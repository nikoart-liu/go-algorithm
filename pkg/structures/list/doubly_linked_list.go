package list

import "fmt"

// DoublyNode represents a node in the doubly linked list.
type DoublyNode[T any] struct {
	Value T
	Next  *DoublyNode[T]
	Prev  *DoublyNode[T]
}

// DoublyLinkedList represents a doubly linked list.
type DoublyLinkedList[T any] struct {
	head *DoublyNode[T]
	tail *DoublyNode[T]
	size int
}

// NewDoubly creates a new empty DoublyLinkedList.
func NewDoubly[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// PushFront adds a value to the front of the list.
func (l *DoublyLinkedList[T]) PushFront(value T) {
	newNode := &DoublyNode[T]{Value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.Next = l.head
		l.head.Prev = newNode
		l.head = newNode
	}
	l.size++
}

// PushBack adds a value to the back of the list.
func (l *DoublyLinkedList[T]) PushBack(value T) {
	newNode := &DoublyNode[T]{Value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.Next = newNode
		newNode.Prev = l.tail
		l.tail = newNode
	}
	l.size++
}

// Size returns the number of elements in the list.
func (l *DoublyLinkedList[T]) Size() int {
	return l.size
}

// Get returns the value at the specified index.
func (l *DoublyLinkedList[T]) Get(index int) (T, error) {
	if index < 0 || index >= l.size {
		var zero T
		return zero, fmt.Errorf("index out of bounds")
	}

	var current *DoublyNode[T]
	if index < l.size/2 {
		current = l.head
		for i := 0; i < index; i++ {
			current = current.Next
		}
	} else {
		current = l.tail
		for i := l.size - 1; i > index; i-- {
			current = current.Prev
		}
	}
	return current.Value, nil
}

// InsertAt inserts a value at the specified index.
func (l *DoublyLinkedList[T]) InsertAt(index int, value T) error {
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

	newNode := &DoublyNode[T]{Value: value}

	var current *DoublyNode[T]
	if index < l.size/2 {
		current = l.head
		for i := 0; i < index; i++ {
			current = current.Next
		}
	} else {
		current = l.tail
		for i := l.size - 1; i > index; i-- {
			current = current.Prev
		}
	}

	prev := current.Prev

	newNode.Next = current
	newNode.Prev = prev

	prev.Next = newNode
	current.Prev = newNode

	l.size++
	return nil
}

// Remove removes the element at the specified index and returns its value.
func (l *DoublyLinkedList[T]) Remove(index int) (T, error) {
	if index < 0 || index >= l.size {
		var zero T
		return zero, fmt.Errorf("index out of bounds")
	}

	var removeNode *DoublyNode[T]

	if index == 0 {
		removeNode = l.head
		l.head = l.head.Next
		if l.head != nil {
			l.head.Prev = nil
		} else {
			l.tail = nil
		}
		l.size--
		return removeNode.Value, nil
	}

	if index == l.size-1 {
		removeNode = l.tail
		l.tail = l.tail.Prev
		if l.tail != nil {
			l.tail.Next = nil
		} else {
			l.head = nil
		}
		l.size--
		return removeNode.Value, nil
	}

	if index < l.size/2 {
		removeNode = l.head
		for i := 0; i < index; i++ {
			removeNode = removeNode.Next
		}
	} else {
		removeNode = l.tail
		for i := l.size - 1; i > index; i-- {
			removeNode = removeNode.Prev
		}
	}

	prev := removeNode.Prev
	next := removeNode.Next

	prev.Next = next
	next.Prev = prev

	l.size--
	return removeNode.Value, nil
}

// Clear removes all elements from the list.
func (l *DoublyLinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}