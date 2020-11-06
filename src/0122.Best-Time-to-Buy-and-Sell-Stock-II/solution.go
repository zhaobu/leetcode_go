package Solution

/*
[解法3 简单的一次遍历](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-ii-by-leetcode/)
*/
func MaxProfit1(prices []int) int {
	sum := 0
	for i := 0; i+1 < len(prices); i++ {
		if prices[i+1] > prices[i] {
			sum += prices[i+1] - prices[i]
		}
	}
	return sum
}

/*
[解法2 峰谷法](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-ii-by-leetcode/)
*/
func MaxProfit2(prices []int) int {
	sum, valley, peak := 0, 0, 0
	for i := 0; i+1 < len(prices); i++ {
		// 找波谷
		for ; i+1 < len(prices) && prices[i] >= prices[i+1]; i++ {
		}
		valley = prices[i]
		// 找波峰
		for ; i+1 < len(prices) && prices[i] <= prices[i+1]; i++ {
		}
		peak = prices[i]
		sum += peak - valley
	}
	return sum
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
		tmp := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, tmp-prices[i])
	}
	return dp_i_0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
