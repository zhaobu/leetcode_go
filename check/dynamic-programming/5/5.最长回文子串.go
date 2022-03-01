package main

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start

/*
暴力法
*/
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start := 0
	maxLen := 0

	var isPalindrome = func(start, end int) bool {
		for ; start < end; start, end = start+1, end-1 {
			if s[start] != s[end] {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if isPalindrome(i, j) && j-i > maxLen {
				start = i
				maxLen = j - i
			}
		}
	}
	return s[start : start+maxLen+1]
}

// @lc code=end
