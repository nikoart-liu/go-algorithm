package dfs

// ---------------------------------------------------------
// 模板 1: 回溯算法 (Backtracking) - 万能公式
// 适用场景: 全排列 (Permutations)、组合 (Combinations)、子集 (Subsets)、N 皇后
// ---------------------------------------------------------

/*
回溯算法核心思想：
result = []
func backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return

    for 选择 in 选择列表:
        1. 做选择 (将该选择从选择列表移除，并加入路径)
        2. backtrack(路径, 选择列表)
        3. 撤销选择 (将该选择恢复到选择列表，并从路径移除)
*/

// Permute 全排列示例模板
func Permute(nums []int) [][]int {
	var res [][]int
	var path []int
	used := make([]bool, len(nums)) // 记录元素是否已被使用

	var backtrack func()
	backtrack = func() {
		// 1. 结束条件: 路径长度等于原数组长度
		if len(path) == len(nums) {
			// 注意：Go 中 slice 是引用，存入结果时必须复制一份，否则会被后续修改
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		// 2. 遍历选择列表
		for i := 0; i < len(nums); i++ {
			// 如果已经用过了，跳过 (剪枝逻辑根据具体题目变)
			if used[i] {
				continue
			}

			// 3. 做选择
			used[i] = true
			path = append(path, nums[i])

			// 4. 进入下一层决策树 (递归)
			backtrack()

			// 5. 撤销选择 (回溯的核心！)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}

// ---------------------------------------------------------
// 模板 2: 树/图的 DFS (Depth First Search)
// 适用场景: 二叉树路径、图的连通性、所有可能的路径
// ---------------------------------------------------------

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func AllPaths(root *TreeNode) [][]int {
	var res [][]int
	var path []int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 1. 处理当前节点 (加入路径)
		path = append(path, node.Val)

		// 2. 叶子节点判断 (结束条件)
		if node.Left == nil && node.Right == nil {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		} else {
			// 3. 递归左右子树
			dfs(node.Left)
			dfs(node.Right)
		}

		// 4. 回溯: 离开当前节点前，将其从路径移除
		path = path[:len(path)-1]
	}

	dfs(root)
	return res
}
