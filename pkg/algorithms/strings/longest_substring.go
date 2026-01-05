package strings

// LengthOfLongestSubstring 解决 LeetCode 第 3 题
// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
func LengthOfLongestSubstring(s string) int {
	// --- 模板开始 ---
	// 1. 初始化窗口和数据记录 (用 map 记录字符出现的次数)
	window := make(map[byte]int)
	left, right := 0, 0
	maxLen := 0

	// 2. 扩大右边界
	for right < len(s) {
		charRight := s[right]
		right++

		// 3. 进窗：增加当前字符计数
		window[charRight]++

		// 4. 收缩判断：如果当前字符出现次数 > 1，说明有重复
		for window[charRight] > 1 {
			// d 是要移出窗口的字符
			charLeft := s[left]
			left++

			// 5. 出窗：减少计数
			window[charLeft]--
		}

		// 6. 更新结果：此时窗口内一定没有重复字符
		// 窗口长度为 right - left
		if right-left > maxLen {
			maxLen = right - left
		}
	}
	// --- 模板结束 ---

	return maxLen
}
