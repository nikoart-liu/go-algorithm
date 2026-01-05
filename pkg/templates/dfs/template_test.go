package dfs

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{1, 2}
	expected := [][]int{
		{1, 2},
		{2, 1},
	}

	got := Permute(nums)

	// 注意：结果的顺序可能不同，这里为了简化测试假设顺序一致
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Permute() = %v, want %v", got, expected)
	}
}

func TestAllPaths(t *testing.T) {
	// 构建二叉树:
	//      1
	//     / \
	//    2   3
	//   /
	//  4
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}

	expected := [][]int{
		{1, 2, 4},
		{1, 3},
	}

	got := AllPaths(root)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("AllPaths() = %v, want %v", got, expected)
	}
}
