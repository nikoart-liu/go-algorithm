package bfs

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	// 构建二叉树:
	//      3
	//     / \
	//    9  20
	//      /  \
	//     15   7
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	expected := [][]int{
		{3},
		{9, 20},
		{15, 7},
	}

	got := LevelOrder(root)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("LevelOrder() = %v, want %v", got, expected)
	}
}

func TestShortestPathInGrid(t *testing.T) {
	// 0: 空地, 1: 障碍物
	// S (0,0) -> E (2,2)
	// 0 0 0
	// 1 1 0
	// 0 0 0
	grid := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{0, 0, 0},
	}

	start := Point{0, 0}
	end := Point{2, 2}

	// 预期路径: (0,0)->(0,1)->(0,2)->(1,2)->(2,2) = 4 步
	expected := 4
	got := ShortestPathInGrid(grid, start, end)

	if got != expected {
		t.Errorf("ShortestPathInGrid() = %v, want %v", got, expected)
	}

	// 测试无法到达的情况
	// 0 1 0
	// 1 1 0
	// 0 0 0
	grid2 := [][]int{
		{0, 1, 0},
		{1, 1, 0},
		{0, 0, 0},
	}
	if got := ShortestPathInGrid(grid2, start, end); got != -1 {
		t.Errorf("ShortestPathInGrid(Unreachable) = %v, want -1", got)
	}
}
