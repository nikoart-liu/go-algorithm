package binarysearch

import (
	"testing"
)

func TestBinarySearchStandard(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"Found middle", []int{1, 2, 3, 4, 5}, 3, 2},
		{"Found first", []int{1, 2, 3, 4, 5}, 1, 0},
		{"Found last", []int{1, 2, 3, 4, 5}, 5, 4},
		{"Not found (too small)", []int{1, 2, 3, 4, 5}, 0, -1},
		{"Not found (too large)", []int{1, 2, 3, 4, 5}, 6, -1},
		{"Empty array", []int{}, 3, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchStandard(tt.nums, tt.target); got != tt.want {
				t.Errorf("BinarySearchStandard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchLowerBound(t *testing.T) {
	// 场景：数组 [1, 2, 2, 2, 5]，找 target = 2
	// 期望返回第一个 2 的索引，即 1
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"Find first occurrence", []int{1, 2, 2, 2, 5}, 2, 1},
		{"Target not found (insertion point)", []int{1, 3, 5}, 2, 1}, // 2 应该插入在 1 和 3 之间，即索引 1
		{"Target greater than all", []int{1, 3, 5}, 6, 3},            // 6 应该插入在末尾，即索引 3 (len)
		{"Target smaller than all", []int{1, 3, 5}, 0, 0},            // 0 应该插入在开头，即索引 0
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchLowerBound(tt.nums, tt.target); got != tt.want {
				t.Errorf("BinarySearchLowerBound() = %v, want %v", got, tt.want)
			}
		})
	}
}
