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