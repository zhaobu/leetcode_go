package main

import "strings"

/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] 删除最外层的括号
 */

// @lc code=start

/*
解法1 记录左括号的个数
*/
func removeOuterParentheses(s string) string {
	//分解出每一个原语
	ret := strings.Builder{}

	left := 0 //左右括号的个数
	for i := range s {
		if s[i] == ')' {
			left--
		}
		if left > 0 {
			ret.WriteByte(s[i])
		}
		if s[i] == '(' {
			left++
		}
	}

	return ret.String()
}

// @lc code=end
