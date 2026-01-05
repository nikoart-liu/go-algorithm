package bfs

// ---------------------------------------------------------
// 基础数据结构定义 (为了演示模板)
// ---------------------------------------------------------

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Point 坐标点 (用于网格搜索)
type Point struct {
	X, Y int
}

// ---------------------------------------------------------
// 模板 1: 二叉树层序遍历 (Tree Level Order Traversal)
// 适用场景: 二叉树的层序遍历、寻找树的最大深度、寻找二叉树每一层的平均值等
// ---------------------------------------------------------
func LevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	// 1. 初始化队列，并将根节点加入队列
	queue := []*TreeNode{root}

	// 2. 循环直到队列为空
	for len(queue) > 0 {
		var currentLevel []int
		levelSize := len(queue) // 记录当前层的节点数 (关键!)

		// 3. 遍历当前层的所有节点
		for i := 0; i < levelSize; i++ {
			// 3.1 出队
			node := queue[0]
			queue = queue[1:]

			// 3.2 处理节点逻辑 (比如记录值)
			currentLevel = append(currentLevel, node.Val)

			// 3.3 将子节点入队 (先左后右)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 4. 将当前层结果存入最终结果
		result = append(result, currentLevel)
	}

	return result
}

// ---------------------------------------------------------
// 模板 2: 网格最短路径 (Grid Shortest Path)
// 适用场景: 迷宫问题、腐烂的橘子、最短步数、骑士最短路径
// ---------------------------------------------------------
func ShortestPathInGrid(grid [][]int, start, end Point) int {
	// 边界检查 (根据具体题目调整)
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	rows, cols := len(grid), len(grid[0])
	
	// 定义四个方向 (上、下、左、右)
	directions := []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// 1. 初始化队列
	queue := []Point{start}

	// 2. 初始化 visited 数组 (非常重要！防止走回头路)
	visited := make(map[Point]bool)
	visited[start] = true

	step := 0 // 记录步数

	// 3. BFS 循环
	for len(queue) > 0 {
		levelSize := len(queue) // 当前层的节点数

		// 遍历当前层
		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:]

			// 判断是否到达终点
			if curr == end {
				return step
			}

			// 4. 探索四个方向
			for _, dir := range directions {
				nextX, nextY := curr.X+dir.X, curr.Y+dir.Y
				next := Point{nextX, nextY}

				// 5. 合法性检查 (越界检查 + 是否访问过 + 题目特定障碍物逻辑)
				if nextX >= 0 && nextX < rows && 
				   nextY >= 0 && nextY < cols && 
				   !visited[next] && 
				   grid[nextX][nextY] != 1 { // 假设 1 是障碍物

					queue = append(queue, next)
					visited[next] = true // 标记为已访问
				}
			}
		}
		step++ // 走完一层，步数 +1
	}

	return -1 // 无法到达
}
