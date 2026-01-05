package twopointers

import (
	"reflect"
	"testing"
)

func TestLeftRightPointers(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	if got := LeftRightPointers(nums, target); !reflect.DeepEqual(got, expected) {
		t.Errorf("LeftRightPointers() = %v, want %v", got, expected)
	}
}

func TestFastSlowPointers(t *testing.T) {
	nums := []int{1, 1, 2, 2, 3}
	// 处理后应为 [1, 2, 3, ...], 长度为 3
	expected := 3
	if got := FastSlowPointers(nums); got != expected {
		t.Errorf("FastSlowPointers() = %v, want %v", got, expected)
	}
}
