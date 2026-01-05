package queue_test

import (
	"testing"
	"go-algorithm/pkg/structures/queue"
)

func BenchmarkSliceQueue_Enqueue(b *testing.B) {
	q := queue.NewSliceQueue[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func BenchmarkListQueue_Enqueue(b *testing.B) {
	q := queue.NewListQueue[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func BenchmarkCircularQueue_Enqueue(b *testing.B) {
	q := queue.NewCircularQueue[int](b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = q.Enqueue(i)
	}
}

func BenchmarkSliceQueue_Dequeue(b *testing.B) {
	q := queue.NewSliceQueue[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = q.Dequeue()
	}
}

func BenchmarkListQueue_Dequeue(b *testing.B) {
	q := queue.NewListQueue[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = q.Dequeue()
	}
}

func BenchmarkCircularQueue_Dequeue(b *testing.B) {
	q := queue.NewCircularQueue[int](b.N)
	for i := 0; i < b.N; i++ {
		_ = q.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = q.Dequeue()
	}
}
