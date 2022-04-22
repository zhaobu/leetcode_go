package main

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start

/*
解法1 模拟相加
*/
func addStrings1(num1 string, num2 string) string {
	pre := 0 //上次相加溢出的部分
	ret := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; {
		cur := pre
		if j >= 0 {
			cur += int(num2[j] - '0')
			j--
		}
		if i >= 0 {
			cur += int(num1[i] - '0')
			i--
		}
		pre = cur / 10
		cur = cur % 10
		ret = strconv.Itoa(cur) + ret
		// fmt.Printf("cur=%v,ret=%+v\n", cur, ret)
	}
	if pre > 0 {
		ret = strconv.Itoa(pre) + ret
	}
	// fmt.Printf("ret=%+v\n", ret)
	return ret
}

// @lc code=end
