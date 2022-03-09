package main

/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 */

// @lc code=start

/*
解法1 动态规划
*/
func wiggleMaxLength1(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	up := make([]int, n)
	down := make([]int, n)

	up[0] = 1
	down[0] = 1
	maxLen := 0
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			up[i] = up[i-1]
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			up[i] = up[i-1]
			down[i] = max(down[i-1], up[i-1]+1)
		} else {
			up[i] = max(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
		}

		maxLen = max(maxLen, max(up[i], down[i]))
	}
	return maxLen
}

/*
解法2 解法1状态压缩
根据状态转移方程优化
*/
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	up := 1
	down := 1
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			down = max(down, up+1)
		} else if nums[i] > nums[i-1] {
			up = max(up, down+1)
		}

	}
	return max(up, down)
}

// @lc code=end
