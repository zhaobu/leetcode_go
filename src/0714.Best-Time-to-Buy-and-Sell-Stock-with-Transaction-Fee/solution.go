package Solution

/*
[labuladong 一个方法团灭 LeetCode 股票买卖问题](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie-qing-an-shun-xu-yue-du/tuan-mie-gu-piao-wen-ti)
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i] - fee)
解释：相当于买入股票的价格升高了。
在第一个式子里减也是一样的，相当于卖出股票的价格减小了。*/
func MaxProfit1(prices []int, fee int) int {
	if len(prices) < 1 {
		return 0
	}
	dp_i0, dp_i1 := 0, -1<<31
	for i := 0; i < len(prices); i++ {
		tmp := dp_i0
		dp_i0 = max(dp_i0, dp_i1+prices[i])
		dp_i1 = max(dp_i1, tmp-prices[i]-fee)
	}
	return dp_i0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
