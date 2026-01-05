package binarysearch

// ---------------------------------------------------------
// 模板 1: 标准二分查找 (Standard Binary Search)
// 适用场景: 在有序数组中查找一个具体的 target，如果不存在返回 -1
// ---------------------------------------------------------
func BinarySearchStandard(nums []int, target int) int {
	// 1. 初始化: 左闭右闭区间 [0, n-1]
	left, right := 0, len(nums)-1

	// 2. 循环条件: left <= right
	// 为什么是 <= ? 因为区间是 [left, right]，当 left==right 时，区间内还有一个元素需要判断
	for left <= right {
		// 3. 计算中点 (防止 (left+right) 溢出)
		mid := left + (right-left)/2

		if nums[mid] == target {
			// 4. 找到目标，直接返回
			return mid
		} else if nums[mid] < target {
			// 目标在右边，搜索区间变为 [mid+1, right]
			left = mid + 1
		} else {
			// 目标在左边，搜索区间变为 [left, mid-1]
			right = mid - 1
		}
	}

	// 5. 未找到
	return -1
}

// ---------------------------------------------------------
// 模板 2: 寻找左侧边界 (Lower Bound) - 最实用的模板！
// 适用场景:
// 1. 数组中有重复元素，找 target 第一次出现的位置
// 2. 找第一个 >= target 的元素位置
// 3. 确定 target 应该插入的位置
// ---------------------------------------------------------
func BinarySearchLowerBound(nums []int, target int) int {
	// 1. 初始化: 左闭右开区间 [0, n)
	// 注意：这里 right = len(nums) 是为了处理 target 比所有元素都大的情况（返回 len）
	left, right := 0, len(nums)

	// 2. 循环条件: left < right
	// 因为是左闭右开 [left, right)，当 left == right 时区间为空，循环结束
	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			// 目标肯定在 mid 右侧，区间变为 [mid+1, right)
			left = mid + 1
		} else {
			// nums[mid] >= target
			// 目标可能就是 mid，或者是 mid 左边的元素
			// 区间缩减为 [left, mid)
			right = mid
		}
	}

	// 3. 返回结果
	// 循环结束时 left == right，指向“第一个 >= target”的位置
	return left
}
