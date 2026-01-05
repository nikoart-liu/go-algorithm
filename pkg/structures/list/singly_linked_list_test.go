package list_test

import (
	"testing"

	"go-algorithm/pkg/structures/list"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_PushFront(t *testing.T) {
	l := list.New[int]()
	assert.Equal(t, 0, l.Size())

	l.PushFront(10)
	assert.Equal(t, 1, l.Size())

	l.PushFront(20)
	assert.Equal(t, 2, l.Size())
}

func TestLinkedList_PushBack(t *testing.T) {
	l := list.New[int]()
	l.PushBack(10)
	assert.Equal(t, 1, l.Size())

	l.PushBack(20)
	assert.Equal(t, 2, l.Size())
}

func TestLinkedList_Size(t *testing.T) {
	l := list.New[string]()
	assert.Equal(t, 0, l.Size())
	l.PushFront("a")
	assert.Equal(t, 1, l.Size())
}

func TestLinkedList_Get(t *testing.T) {
	l := list.New[int]()
	l.PushBack(10)
	l.PushBack(20)

	val, err := l.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 10, val)

	val, err = l.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, 20, val)

	_, err = l.Get(-1)
	assert.Error(t, err)

	_, err = l.Get(2)
	assert.Error(t, err)
}

func TestLinkedList_InsertAt(t *testing.T) {
	l := list.New[int]()
	// Insert at 0 (empty)
	err := l.InsertAt(0, 10)
	assert.NoError(t, err)
	assert.Equal(t, 1, l.Size())
	val, _ := l.Get(0)
	assert.Equal(t, 10, val)

	// Insert at 1 (end)
	err = l.InsertAt(1, 30)
	assert.NoError(t, err)
	assert.Equal(t, 2, l.Size())
	val, _ = l.Get(1)
	assert.Equal(t, 30, val)

	// Insert at 1 (middle) -> 10, 20, 30
	err = l.InsertAt(1, 20)
	assert.NoError(t, err)
	assert.Equal(t, 3, l.Size())
	val, _ = l.Get(1)
	assert.Equal(t, 20, val)
	val, _ = l.Get(2)
	assert.Equal(t, 30, val)

	// Invalid index
	err = l.InsertAt(5, 50)
	assert.Error(t, err)
}

func TestLinkedList_Remove(t *testing.T) {
	l := list.New[int]()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)

	// Remove middle (index 1, value 20)
	val, err := l.Remove(1)
	assert.NoError(t, err)
	assert.Equal(t, 20, val)
	assert.Equal(t, 2, l.Size())

	// Check remaining: 10, 30
	v0, _ := l.Get(0)
	assert.Equal(t, 10, v0)
	v1, _ := l.Get(1)
	assert.Equal(t, 30, v1)

	// Remove head (10)
	val, err = l.Remove(0)
	assert.NoError(t, err)
	assert.Equal(t, 10, val)

	// Remove tail (30) - which is now at 0
	val, err = l.Remove(0)
	assert.NoError(t, err)
	assert.Equal(t, 30, val)
	assert.Equal(t, 0, l.Size())

	// Remove from empty
	_, err = l.Remove(0)
	assert.Error(t, err)
}

func TestLinkedList_Clear(t *testing.T) {
	l := list.New[int]()
	l.PushBack(1)
	l.Clear()
	assert.Equal(t, 0, l.Size())
	_, err := l.Get(0)
	assert.Error(t, err)
}
