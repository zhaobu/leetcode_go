package main

/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start

/*
解法1: 动态规划

状态定义: 第i天结束时,只会有持有和不持有股票2种状态
dp[i][0]表示第i天结束后不持有股票能获得的最大利润
dp[i][1]表示第i天结束后持有股票能获得的最大利润

basecase:
dp[0][0] = 0          第0天结束后不持有股票
dp[0][1] = -prices[0] 第0天结束后持有股票

状态转移方程:
第i天不持有股票,有3种情况
1. 第 i-1 天本来就不持有股票,第 i 天也没进行任何操作 dp[i][0] = dp[i-1][0]
2. 第 i-1 天本来就不持有股票,第 i 天买了股票又卖出 dp[i][0] = dp[i-1][0]-prices[i]+prices[i]
3. 第 i-1 天持有股票,第 i-1 天结束时股票收益是 dp[i-1][1],外加手里还有一支股票,然后第i天卖了
就能赚 prices[i]. 所以 dp[i][0] = dp[i-1][1]+price[i]
综合来说 dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])

第i天持有股票,也有3种情况
1. 第 i-1 天持已经持有股票,并且第i天也没操作. 所以第i天收益相对第 i-1 天不变 dp[i][1] = dp[i-1][1]
2. 第 i-1 天不持有股票,第 i 天新买进了股票. dp[i][1] = dp[i-1][0] - price[i]
3. 同理也可能存在第 i-1 天持已经持有股票,并且第i天先卖出又买进 dp[i][1] = dp[i-1][1]+prices[i]-prices[i]

综合得出转移方程
dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
*/
func maxProfit1(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	/*
	   dp[i][0]表示第i天结束后不持有股票能获得的最大利润
	   dp[i][1]表示第i天结束后持有股票能获得的最大利润
	*/
	dp := make([][2]int, len(prices))
	dp[0][0] = 0          //第0天结束后不持有股票
	dp[0][1] = -prices[0] //第0天结束后持有股票

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[len(prices)-1][0]
}

/*
解法2 动态规划 优化空间复杂度
观察解法1的状态转移方程,只和dp[i-1]相关,所以完全可以只用一维的[2]int来代替
*/

func maxProfit(prices []int) int {
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
		dp0 := dp[0]
		dp[0] = max(dp[0], dp[1]+prices[i])
		dp[1] = max(dp[1], dp0-prices[i])
	}

	return dp[0]
}

/*
解法3 贪心
因为操作次数没限制,只要不是同时买入多股,所以只需要
买入所有股票上升的区间就是能获得的最高利润

*/
func maxProfit3(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	sum := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] >= prices[i-1] {
			sum += prices[i] - prices[i-1]
		}
	}

	return sum
}

// @lc code=end
