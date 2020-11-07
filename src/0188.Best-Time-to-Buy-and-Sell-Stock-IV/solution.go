package Solution

/*
[labuladong 一个方法团灭 LeetCode 股票买卖问题](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie-qing-an-shun-xu-yue-du/tuan-mie-gu-piao-wen-ti)
*/
func MaxProfit1(k int, prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	dp := make([][][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = append(dp[i], make([][2]int, k+1)...)
		//穷举所有可能的k,k表示进行交易的次数
		for k := k; k >= 1; k-- {
			if i-1 == -1 {
				dp[0][k][0] = 0
				dp[0][k][1] = -prices[0]
			} else {
				dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
				dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
			}
		}
	}
	return dp[len(prices)-1][k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
