package Solution

/*
[labuladong 一个方法团灭 LeetCode 股票买卖问题](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie-qing-an-shun-xu-yue-du/tuan-mie-gu-piao-wen-ti)
*/
func maxProfit1(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	/*
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])
		解释：第 i 天选择 buy 的时候，要从 i-2 的状态转移，而不是 i-1 。
	*/
	dp_i_0, dp_i_1 := 0, -1<<31
	dp_pre_0 := 0 // 代表 dp[i-2][0]
	for i := 0; i < len(prices); i++ {
		temp := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, dp_pre_0-prices[i])
		dp_pre_0 = temp
	}
	return dp_i_0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
