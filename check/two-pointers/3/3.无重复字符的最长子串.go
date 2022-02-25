package main

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

/*
 解法1
 + . 从做往右依次求以每个字符结尾的最长无重复字符子串的长度,并和记录的最大值比较更新
 + . lastCharIdx 表示当前字符最近出现的位置
 + . 用一个map charMap记录已经遍历过的位置所有字符的最近出现的下标
 + . 用一个变量 preStartIdx 记录以上一个字符结尾的最长无重复字符子串的开始下标
 进行如下比较:
 1. 如果lastCharIdx 不存在或者 lastCharIdx < preStartIdx 说明可以直接把当前字符拼接在上一个
 字符结尾的最长无重复字符子串后面
 2. 如果不满足1说明以当前字符结尾的最长无重复字符子串的开始位置只能是 lastCharIdx+1

*/
// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	/*
		"abcabcbb"
	*/

	charMap := map[byte]int{} //每个字符最近出现的位置
	lastCharIdx := 0          //当前字符最近出现的位置
	preStartIdx := 0          //以上一个位置所在字符结尾的最长子串的开始位置
	curLen := 0               //以当前字符结尾的最长子串长度
	maxLen := 0               //无重复子串的最长长度
	ok := false               //当前字符是否在之前出现过
	for i := 0; i < len(s); i++ {
		lastCharIdx, ok = charMap[s[i]]
		if !ok || lastCharIdx < preStartIdx {
			curLen = i - preStartIdx
		} else {
			curLen = i - (lastCharIdx + 1)
			preStartIdx = lastCharIdx + 1
		}
		if curLen > maxLen {
			maxLen = curLen
		}
		charMap[s[i]] = i
	}
	return maxLen + 1
}

// @lc code=end
