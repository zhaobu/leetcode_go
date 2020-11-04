package Solution

/*
[labuladong 滑动窗口算法 专题](https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485141&idx=1&sn=0e4583ad935e76e9a3f6793792e60734&chksm=9bd7f8ddaca071cbb7570b2433290e5e2628d20473022a5517271de6d6e50783961bebc3dd3b&scene=21#wechat_redirect)
*/
func lengthOfLongestSubstring1(s string) int {
	window := make(map[byte]int)
	left, right, res := 0, 0, 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}

// [leetcode题解 ](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/longest-substring-without-repeating-characters-b-2/)
func lengthOfLongestSubstring2(s string) int {
	//记录一个字符如果后面出现重复时,start应该调整到的新位置
	m := make(map[rune]int)
	//start表示子串的起始位置,max记录不包含重复字符的子串的最大长度
	start, max := -1, 0
	//依次遍历s
	for k, v := range s {
		if last, ok := m[v]; ok && last > start { //如果发现重复字符
			start = last
		}
		m[v] = k
		if k-start > max { //保存最大值
			max = k - start
		}
	}
	return max
}
