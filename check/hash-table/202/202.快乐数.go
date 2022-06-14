package main

/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start

/*
解法1
*/
func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	exist := map[int]bool{}

	cur := n
	for cur != 1 {
		//
		next := 0
		for cur != 0 {
			last := cur % 10
			next += last * last
			cur = cur / 10
		}
		if exist[next] {
			return false
		}
		exist[next] = true
		cur = next
	}
	return true
}

// @lc code=end
