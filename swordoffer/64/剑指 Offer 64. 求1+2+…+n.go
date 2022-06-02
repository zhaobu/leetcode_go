package main

import (
	. "leetcode/check"
)

/*
 *
 *
 * 剑指 Offer 64. 求1+2+…+n
 */

// @lc code=start

/*
 解法1
*/
func sumNums(n int) int {
	sum := 0
	var dfs func(n int) bool
	dfs = func(n int) bool {
		sum += n
		return n > 0 && dfs(n-1)
	}
	dfs(n)
	return sum
}

// @lc code=end
