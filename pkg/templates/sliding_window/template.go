package slidingwindow

// ---------------------------------------------------------
// 模板: 通用滑动窗口 (Sliding Window)
// 适用场景: 最小覆盖子串、最长无重复子串、子数组最大和
// ---------------------------------------------------------
func SlidingWindow(s string) int {
	// 1. 初始化窗口、数据记录 (如 Map) 和结果
	window := make(map[byte]int)
	left, right := 0, 0
	res := 0

	// 2. 扩大右边界
	for right < len(s) {
		// c 是将移入窗口的字符
		charRight := s[right]
		right++

		// 3. 进窗：进行窗口内数据的一系列更新
		window[charRight]++

		// 4. 判断左侧窗口是否要收缩 (根据题目要求：比如长度超了、出现重复了)
		for window[charRight] > 1 { // 示例：最长无重复子串的判断条件
			// d 是将移出窗口的字符
			charLeft := s[left]
			left++

			// 5. 出窗：进行窗口内数据的一系列更新
			window[charLeft]--
		}

		// 6. 更新最终结果 (在窗口收缩完后更新，此时窗口一定满足条件)
		if right-left > res {
			res = right - left
		}
	}
	return res
}
