package queue_test

import (
	"testing"

	"go-algorithm/pkg/structures/queue"

	"github.com/stretchr/testify/assert"
)

func TestSliceQueue(t *testing.T) {
	q := queue.NewSliceQueue[int]()
	assert.True(t, q.IsEmpty())

	q.Enqueue(10)
	q.Enqueue(20)
	assert.Equal(t, 2, q.Size())

	val, err := q.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 10, val)

	val, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 10, val)
	assert.Equal(t, 1, q.Size())

	val, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 20, val)
	assert.True(t, q.IsEmpty())

	_, err = q.Dequeue()
	assert.Error(t, err)

	_, err = q.Peek()
	assert.Error(t, err)
}

func TestListQueue(t *testing.T) {
	q := queue.NewListQueue[int]()
	assert.True(t, q.IsEmpty())

	q.Enqueue(10)
	q.Enqueue(20)
	assert.Equal(t, 2, q.Size())

	val, err := q.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 10, val)

	val, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 10, val)

	val, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 20, val)
	assert.True(t, q.IsEmpty())

	_, err = q.Dequeue()
	assert.Error(t, err)

	_, err = q.Peek()
	assert.Error(t, err)
}

func TestCircularQueue(t *testing.T) {
	q := queue.NewCircularQueue[int](3)
	assert.True(t, q.IsEmpty())
	assert.False(t, q.IsFull())

	// Enqueue 1, 2, 3
	assert.NoError(t, q.Enqueue(10))
	assert.NoError(t, q.Enqueue(20))
	assert.NoError(t, q.Enqueue(30))
	assert.True(t, q.IsFull())
	assert.Equal(t, 3, q.Size())

	// Full enqueue error
	assert.Error(t, q.Enqueue(40))

	// Dequeue 1
	val, err := q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 10, val)
	assert.False(t, q.IsFull())

	// Enqueue 4 (wrap around)
	assert.NoError(t, q.Enqueue(40))
	assert.True(t, q.IsFull())

	// Dequeue all
	v2, _ := q.Dequeue()
	assert.Equal(t, 20, v2)
	v3, _ := q.Dequeue()
	assert.Equal(t, 30, v3)
	v4, _ := q.Dequeue()
	assert.Equal(t, 40, v4)
	assert.True(t, q.IsEmpty())

	// Empty dequeue error
	_, err = q.Dequeue()
	assert.Error(t, err)
}