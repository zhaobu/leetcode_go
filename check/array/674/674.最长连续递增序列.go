package main

/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */

// @lc code=start

/*
解法1 动态规划 空间复杂度优化版
*/
func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	/*
		1.dp[i] 表示以nums[i]结尾的最长连续递增子序列的长度
		2. 当 i=0 时递增子序列就是nums[0]
		3. 因为dp[i]只和dp[i-1]相关. 可以优化空间复杂度
	*/
	dp := 1
	maxLen := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp += 1
		} else {
			dp = 1
		}
		if dp > maxLen {
			maxLen = dp
		}
	}

	return maxLen
}

// @lc code=end
