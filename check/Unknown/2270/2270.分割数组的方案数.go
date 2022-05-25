package main

/*
 * @lc app=leetcode.cn id=2270 lang=golang
 *
 * [2270] 分割数组的方案数
 */

// @lc code=start

/*
解法1 前缀和
*/
func waysToSplitArray1(nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1) //用来保存前缀和
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	ret := 0
	for i := 0; i < n-1; i++ {
		if preSum[i+1] >= preSum[n]-preSum[i+1] {
			ret++
		}
	}

	return ret
}

/*
解法2 前缀和(优化空间复杂度)
不提前计算前缀和,而是实时维护
*/
func waysToSplitArray(nums []int) int {
	right := 0
	for _, v := range nums {
		right += v
	}

	ret, left := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		left += nums[i]
		right -= nums[i]
		if left >= right {
			ret++
		}
	}

	return ret
}

// @lc code=end
