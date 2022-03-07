package main

/*
 * @lc app=leetcode.cn id=673 lang=golang
 *
 * [673] 最长递增子序列的个数
 */

// @lc code=start
func findNumberOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	// dp[i]表示nums[i]结尾的最长递增子序列长度
	dp := make([]int, len(nums))
	dp[0] = 1

	maxLen := 1
	sameMaxLenNum := 1
	maxLenNum := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				dpNew := dp[j] + 1
				if dpNew < dp[i] {
					continue
				}
				dp[i] = dpNew
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			sameMaxLenNum = 1
		} else if dp[i] == maxLen {
			sameMaxLenNum++
		}
		if sameMaxLenNum > maxLenNum {
			maxLenNum = sameMaxLenNum
		}
	}

	return maxLenNum
}

// @lc code=end
