package stack_test

import (
	"testing"

	"go-algorithm/pkg/structures/stack"

	"github.com/stretchr/testify/assert"
)

func TestSliceStack(t *testing.T) {
	s := stack.NewSliceStack[int]()
	assert.True(t, s.IsEmpty())

	s.Push(10)
	s.Push(20)
	assert.Equal(t, 2, s.Size())
	assert.False(t, s.IsEmpty())

	val, err := s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 20, val)

	val, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 20, val)
	assert.Equal(t, 1, s.Size())

	val, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 10, val)
	assert.True(t, s.IsEmpty())

	_, err = s.Pop()
	assert.Error(t, err)

	_, err = s.Peek()
	assert.Error(t, err)
}

func TestListStack(t *testing.T) {
	s := stack.NewListStack[int]()
	assert.True(t, s.IsEmpty())

	s.Push(10)
	s.Push(20)
	assert.Equal(t, 2, s.Size())

	val, err := s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 20, val)

	val, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 20, val)

	val, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 10, val)
	assert.True(t, s.IsEmpty())

	_, err = s.Pop()
	assert.Error(t, err)

	_, err = s.Peek()
	assert.Error(t, err)
}