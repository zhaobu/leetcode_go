package check

import "strings"

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	var (
		i = 0
		j = len(s) - 1
	)
	for i < j {
		for ; i < j && !isChar(s[i]); i++ {
		}
		for ; i < j && !isChar(s[j]); j-- {
		}
		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func isChar(ch byte) bool {
	return (ch <= 'Z' && ch >= 'A') || (ch <= 'z' && ch >= 'a') || (ch <= '9' && ch >= '0')
}

// @lc code=end
