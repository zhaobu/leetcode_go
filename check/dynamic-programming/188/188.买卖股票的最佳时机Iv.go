package main

import "math"

/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start

/*
解法1 动态规划

*/
func maxProfit(k int, prices []int) int {
	if k < 1 || len(prices) < 2 {
		return 0
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(prices)
	k = k + 1
	dp := make([][][2]int, n)
	for i, _ := range dp {
		dp[i] = make([][2]int, k)
		dp[i][0][0] = 0
		dp[i][0][1] = math.MinInt

	}

	for j := 0; j < k; j++ {
		dp[0][j][1] = -prices[0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}

	return dp[n-1][k-1][0]
}

// @lc code=end
