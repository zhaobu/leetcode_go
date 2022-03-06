package main

/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start

/*
解法1 动态规划
参考 [买卖股票的最佳时机 II] (https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)
唯一不同就是每次卖出是减去手续费
*/
func maxProfit(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	/*
	   dp[i][0]表示第i天结束后不持有股票能获得的最大利润
	   dp[i][1]表示第i天结束后持有股票能获得的最大利润
	*/
	dp := make([]int, len(prices))
	dp[0] = 0          //第0天结束后不持有股票
	dp[1] = -prices[0] //第0天结束后持有股票

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(prices); i++ {
		dp[0] = max(dp[0], dp[1]+prices[i]-fee)
		dp[1] = max(dp[1], dp[0]-prices[i])
	}

	return dp[0]
}

// @lc code=end
