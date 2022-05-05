package main

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start

/*
解法1 动态规划
*/
func maxProduct1(nums []int) int {
	m := len(nums)
	if m == 1 {
		return nums[0]
	}
	dpMin := make([]int, m) //dpMin[i]表示以nums[i]结尾的连续子数组的最小乘积
	dpMax := make([]int, m) //dpMax[i]表示以nums[i]结尾的连续子数组的最大乘积
	dpMax[0] = nums[0]
	dpMin[0] = nums[0]

	ret := dpMax[0]

	max := func(a, b, c int) int {
		if a > b {
			if a > c {
				return a
			}
			return c
		} else {
			if b > c {
				return b
			}
			return c
		}
	}
	min := func(a, b, c int) int {
		if a < b {
			if a < c {
				return a
			}
			return c
		} else {
			if b < c {
				return b
			}
			return c
		}
	}

	for i := 1; i < m; i++ {
		dpMax[i] = max(nums[i], nums[i]*dpMax[i-1], nums[i]*dpMin[i-1])
		dpMin[i] = min(nums[i], nums[i]*dpMax[i-1], nums[i]*dpMin[i-1])
		if dpMax[i] > ret {
			ret = dpMax[i]
		}
		// fmt.Printf("cur=%+v,dp[i]=%+v,ret=%+v \n", cur, dp[i], ret)
	}

	return ret
}

/*
解法2 动态规划
解法1的空间优化
*/
func maxProduct(nums []int) int {
	m := len(nums)
	if m == 1 {
		return nums[0]
	}
	dpMin, dpMax := nums[0], nums[0]
	ret := dpMax
	max := func(a, b, c int) int {
		if a > b {
			if a > c {
				return a
			}
			return c
		} else {
			if b > c {
				return b
			}
			return c
		}
	}
	min := func(a, b, c int) int {
		if a < b {
			if a < c {
				return a
			}
			return c
		} else {
			if b < c {
				return b
			}
			return c
		}
	}

	for i := 1; i < m; i++ {
		oldMin, oldMax := dpMin, dpMax
		dpMax = max(nums[i], nums[i]*oldMax, nums[i]*oldMin)
		dpMin = min(nums[i], nums[i]*oldMax, nums[i]*oldMin)
		if dpMax > ret {
			ret = dpMax
		}
		// fmt.Printf("cur=%+v,dp[i]=%+v,ret=%+v \n", cur, dp[i], ret)
	}

	return ret
}

// @lc code=end
