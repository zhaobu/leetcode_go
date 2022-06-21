package main

import "strings"

/*
 * @lc app=leetcode.cn id=1108 lang=golang
 *
 * [1108] IP 地址无效化
 */

// @lc code=start
func defangIPaddr(address string) string {
	return strings.Join(strings.Split(address, "."), "[.]")
}

// @lc code=end
