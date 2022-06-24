package main

import "strconv"

/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	n1, n2 := len(a), len(b)

	i, j := n1-1, n2-1
	pre := 0
	ret := ""

	for i >= 0 || j >= 0 {
		cur := pre
		if i >= 0 {
			cur += int(a[i] - '0')
		}
		if j >= 0 {
			cur += int(b[j] - '0')
		}
		// fmt.Printf("cur=%d\n", cur)
		if cur == 3 {
			ret = "1" + ret
			pre = 1
		} else if cur == 2 {
			ret = "0" + ret
			pre = 1
		} else {
			ret = strconv.Itoa(cur) + ret
			pre = 0
		}
		i, j = i-1, j-1
	}
	if pre == 1 {
		ret = "1" + ret
	}
	return ret
}

// @lc code=end
