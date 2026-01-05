package stack_test

import (
	"testing"
	"go-algorithm/pkg/structures/stack"
)

func BenchmarkSliceStack_Push(b *testing.B) {
	s := stack.NewSliceStack[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkListStack_Push(b *testing.B) {
	s := stack.NewListStack[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkSliceStack_Pop(b *testing.B) {
	s := stack.NewSliceStack[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = s.Pop()
	}
}

func BenchmarkListStack_Pop(b *testing.B) {
	s := stack.NewListStack[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = s.Pop()
	}
}
