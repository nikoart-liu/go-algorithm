package list_test

import (
	"testing"

	"go-algorithm/pkg/structures/list"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList_PushFront(t *testing.T) {
	l := list.NewDoubly[int]()
	assert.Equal(t, 0, l.Size())

	l.PushFront(10)
	assert.Equal(t, 1, l.Size())

	l.PushFront(20)
	assert.Equal(t, 2, l.Size())

	val, _ := l.Get(0)
	assert.Equal(t, 20, val)
	val, _ = l.Get(1)
	assert.Equal(t, 10, val)
}

func TestDoublyLinkedList_PushBack(t *testing.T) {
	l := list.NewDoubly[int]()
	l.PushBack(10)
	assert.Equal(t, 1, l.Size())

	l.PushBack(20)
	assert.Equal(t, 2, l.Size())

	val, _ := l.Get(0)
	assert.Equal(t, 10, val)
	val, _ = l.Get(1)
	assert.Equal(t, 20, val)
}

func TestDoublyLinkedList_InsertAt(t *testing.T) {
	l := list.NewDoubly[int]()
	l.InsertAt(0, 10) // PushFront
	l.InsertAt(1, 30) // PushBack
	l.InsertAt(1, 20) // Middle

	// 10, 20, 30
	assert.Equal(t, 3, l.Size())
	v0, _ := l.Get(0)
	assert.Equal(t, 10, v0)
	v1, _ := l.Get(1)
	assert.Equal(t, 20, v1)
	v2, _ := l.Get(2)
	assert.Equal(t, 30, v2)
}

func TestDoublyLinkedList_Remove(t *testing.T) {
	l := list.NewDoubly[int]()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)

	// Remove middle
	val, err := l.Remove(1)
	assert.NoError(t, err)
	assert.Equal(t, 20, val)

	// 10, 30
	assert.Equal(t, 2, l.Size())
	v0, _ := l.Get(0)
	assert.Equal(t, 10, v0)
	v1, _ := l.Get(1)
	assert.Equal(t, 30, v1)

	// Remove head
	l.Remove(0)
	v0, _ = l.Get(0)
	assert.Equal(t, 30, v0)

	// Remove tail
	l.Remove(0)
	assert.Equal(t, 0, l.Size())
}

func TestDoublyLinkedList_Clear(t *testing.T) {
	l := list.NewDoubly[int]()
	l.PushBack(1)
	l.Clear()
	assert.Equal(t, 0, l.Size())
}

func TestDoublyLinkedList_Coverage(t *testing.T) {
	l := list.NewDoubly[int]()
	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}

	// Get from second half
	val, err := l.Get(8)
	assert.NoError(t, err)
	assert.Equal(t, 8, val)

	// Insert in second half
	err = l.InsertAt(8, 88)
	assert.NoError(t, err)
	v8, _ := l.Get(8)
	assert.Equal(t, 88, v8)
	v9, _ := l.Get(9)
	assert.Equal(t, 8, v9) // Shifted

	// Remove from second half
	val, err = l.Remove(8)
	assert.NoError(t, err)
	assert.Equal(t, 88, val)

	// Check tail removal from large list
	l.Remove(l.Size() - 1)
	assert.Equal(t, 9, l.Size())

	// Errors
	_, err = l.Get(-1)
	assert.Error(t, err)
	_, err = l.Get(100)
	assert.Error(t, err)

	err = l.InsertAt(-1, 0)
	assert.Error(t, err)
	err = l.InsertAt(100, 0)
	assert.Error(t, err)

	_, err = l.Remove(-1)
	assert.Error(t, err)
	_, err = l.Remove(100)
	assert.Error(t, err)
}
