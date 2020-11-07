package Solution

/*
[labuladong 一个方法团灭 LeetCode 股票买卖问题](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie-qing-an-shun-xu-yue-du/tuan-mie-gu-piao-wen-ti)
*/
func MaxProfit1(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	max_k := 2
	dp := make([][2 + 1][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		//穷举所有可能的k,k表示进行交易的次数
		for k := max_k; k >= 1; k-- {
			if i-1 == -1 {
				dp[0][k][0] = 0
				dp[0][k][1] = -prices[0]
			} else {
				dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
				dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
			}
		}
	}
	return dp[len(prices)-1][max_k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
[labuladong 一个方法团灭 LeetCode 股票买卖问题,优化空间复杂度 ](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie-qing-an-shun-xu-yue-du/tuan-mie-gu-piao-wen-ti)
*/
func MaxProfit2(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	/*
	   dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1] + prices[i])
	   dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0] - prices[i])
	   dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
	   dp[i][1][1] = max(dp[i-1][1][1], -prices[i])
	*/

	dp_i10, dp_i11 := 0, -1<<31
	dp_i20, dp_i21 := 0, -1<<31
	for i := 0; i < len(prices); i++ {
		dp_i20 = max(dp_i20, dp_i21+prices[i])
		dp_i21 = max(dp_i21, dp_i10-prices[i])

		dp_i10 = max(dp_i10, dp_i11+prices[i])
		// dp_i11 = max(dp_i11, dp_i01-prices[i])
		dp_i11 = max(dp_i11, -prices[i])
	}
	return dp_i20
}
