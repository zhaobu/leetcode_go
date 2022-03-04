package main

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start

/*
解法1 动态规划
状态定义: dp[i] 表示以下标i结尾的最长子序列的长度
basecase:
dp[0] = 1 表示以nums[0] 结尾的最长子序列就是nums[0]
状态转移方程:
在求dp[i]时,遍历 j ∈ [0,i)
1. 当 nums[j] > nums[i] 时,跳过
2. 当 nums[j] <= nums[i] 时说明nums[i]可以接在nums[j]后面
dp[i] = dp[j] + 1
*/
func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	//dp[i] 表示以下标i结尾的最长子序列的长度
	dp := make([]int, len(nums))
	dp[0] = 1
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1 //默认初始化为 最长上升子序列只有nums[i]
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			if dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
		// fmt.Printf("dp[%d]=%+v\n", i, dp[i])
	}

	return maxLen
}

// @lc code=end
