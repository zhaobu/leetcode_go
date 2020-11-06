package Solution

/*
[leetcode 官方 方法二：一次遍历](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/gu-piao-wen-ti-python3-c-by-z1m/)
*/
func MaxProfit1(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	minPrice := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}
	return profit
}

/*
[leetcode题解 动态规划](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/gu-piao-wen-ti-python3-c-by-z1m/)
*/
func MaxProfit2(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	//dp[i]表示前i天的最大利润
	dp := make([]int, len(prices))
	dp[0] = 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		dp[i] = max(dp[i-1], prices[i]-minPrice)
	}
	return dp[len(prices)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
[labuladong 一个方法团灭 LeetCode 股票买卖问题](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie-qing-an-shun-xu-yue-du/tuan-mie-gu-piao-wen-ti)
*/
func MaxProfit3(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	dp_i_0, dp_i_1 := 0, -1<<31
	for i := 0; i < len(prices); i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, -prices[i])
	}
	return dp_i_0
}
