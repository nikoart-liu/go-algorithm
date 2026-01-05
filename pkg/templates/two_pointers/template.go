package twopointers

// ---------------------------------------------------------
// 模板 1: 左右指针 (Left/Right Pointers)
// 适用场景: 有序数组找两数之和、反转数组、回文串判断
// ---------------------------------------------------------
func LeftRightPointers(nums []int, target int) []int {
	// 1. 初始化左右指针
	left, right := 0, len(nums)-1

	// 2. 向中间靠拢
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			// 3. 根据逻辑移动指针
			left++
		} else {
			right--
		}
	}
	return nil
}

// ---------------------------------------------------------
// 模板 2: 快慢指针 (Fast/Slow Pointers)
// 适用场景: 链表找中点、判断环、删除数组重复项
// ---------------------------------------------------------
func FastSlowPointers(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 1. 初始化快慢指针
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		// 2. 当快指针满足某种条件时，移动慢指针
		if nums[fast] != nums[slow] {
			slow++
			// 3. 进行操作 (比如原地修改)
			nums[slow] = nums[fast]
		}
	}
	// 返回处理后的长度
	return slow + 1
}
