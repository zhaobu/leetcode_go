package main

import "strings"

/*
 * @lc app=leetcode.cn id=1108 lang=golang
 *
 * [1108] IP 地址无效化
 */

// @lc code=start

/*
解法1 库函数
*/
func defangIPaddr1(address string) string {
	return strings.Join(strings.Split(address, "."), "[.]")
}

/*
解法2 一次遍历
*/
func defangIPaddr(address string) string {
	strBuid := strings.Builder{}
	start := 0
	for i, v := range address {
		if v == '.' {
			strBuid.WriteString(address[start:i])
			strBuid.WriteString("[.]")
			start = i + 1
		}
	}
	strBuid.WriteString(address[start:])
	return strBuid.String()
}

// @lc code=end
