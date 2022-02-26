package main

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
// @lc code=start

/*
解法3
1. 从左往右遍历,i为无重复字符串的开始位置,j表示该字符串的结束位置
2. 枚举以每个s[i]开始的字符串,找到合适的j的位置
3. 第 i+1个字符作为起始位置时，首先从 i+1 到j 位置都是不重复的,所以从j的下一个位置
 开始判断
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	maxLen := 0
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	charMap := map[byte]int{}
	for i, j := 0, -1; i < len(s); i++ {
		for ; j+1 < len(s); j++ {
			if _, ok := charMap[s[j+1]]; ok {
				break
			}
			charMap[s[j+1]] = j + 1
		}
		maxLen = max(maxLen, j-i+1)
		delete(charMap, s[i])
	}
	return maxLen
}

/*
 解法2
1.从左往右遍历字符串
2.2个变量j和i,j表示以字符i结尾的最长无重复字符串的开始位置,
charMap表示s[i]之前出现过的位置
3.i从左向右遍历,如果s[i]之前没出现过,i就会一直向右遍历,j始终指向不重复字符子串的开始
 如果遇到重复的,说明charMap[i]之前有值,说明j要指向s[i]之前出现的位置的下一个位置也就是
 charMap[s[i]]+1
*/
func lengthOfLongestSubstring2(s string) int {
	if len(s) < 1 {
		return 0
	}
	charMap := map[byte]int{} //记录已遍历的每个字符最近出现的位置
	maxLen := 0               //记录找到的最长无重复字符子串的长度
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	/*
	 abcabcbb
	 j表示以当前字符结尾的最长无重复字符子串的开始位置
	*/
	for i, j := 0, 0; i < len(s); i++ {
		if _, ok := charMap[s[i]]; ok {
			j = max(j, charMap[s[i]]+1)
		}
		charMap[s[i]] = i
		maxLen = max(maxLen, i-j+1)
	}

	return maxLen
}

/*3
解法1
+ 从做往右依次求以每个字符结尾的最长无重复字符子串的长度,并和记录的最大值比较更新
+ lastCharIdx 表示当前字符最近出现的位置
+ 用一个map charMap记录已经遍历过的位置所有字符的最近出现的下标
+ 用一个变量 preStartIdx 记录以上一个字符结尾的最长无重复字符子串的开始下标
进行如下比较:
1. 如果lastCharIdx 不存在或者 lastCharIdx < preStartIdx 说明可以直接把当前字符拼接在上一个
字符结尾的最长无重复字符子串后面
2. 如果不满足1说明以当前字符结尾的最长无重复字符子串的开始位置只能是 lastCharIdx+1

*/
// @lc code=start
func lengthOfLongestSubstring1(s string) int {
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
