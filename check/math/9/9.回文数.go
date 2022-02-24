package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	return true
}

// @lc code=end

/*
题目:[面试题 01.09. 字符串轮转](https://leetcode-cn.com/problems/string-rotation-lcci/)
示例1:

 输入：s1 = "waterbottle", s2 = "erbottlewat"
 输出：True
"waterbottlewaterbottle"
示例2:

 输入：s1 = "aa", s2 = "aba"
 输出：False

*/
func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	s := fmt.Sprintf("%s%s", s1, s1)
	return strings.Contains(s, s2)
}
