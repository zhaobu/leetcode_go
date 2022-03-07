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
状态定义:
1. dp[i][k][0 or 1]
0 <= i <= n - 1, 1 <= k <= K
n 为天数，大 K 为交易数的上限，0 和 1 代表是否持有股票。
比如说 dp[3][2][1] 的含义就是：今天是第三天，我现在手上持有着股票，至今最多进行 2 次交易。
再比如 dp[2][3][0] 的含义：今天是第二天，我现在手上没有持有股票，至今最多进行 3 次交易
2. 最终要求的答案是 dp[n - 1][K][0]因为 dp[n - 1][K][1] 代表到最后一天手上还持有股票，
dp[n - 1][K][0] 表示最后一天手上的股票已经卖出去了，很显然后者得到的利润一定大于前者

basecase:
dp[0][j][0] = 0
解释：第0天不持有股票,不管交易次数限制多少次,同一天内交易与否,利润都为0

dp[0][j][1] = -prices[0]
解释：第0天持有股票,不管交易次数限制多少,最终利润都为-prices[0]

dp[i][0][0] = 0
解释：因为 k 是从 1 开始的，所以 k = 0 意味着根本不允许交易，这时候利润当然是 0。

dp[i][0][1] = -infinity
解释：不允许交易的情况下，是不可能持有股票的。
因为我们的算法要求一个最大值，所以初始值设为一个最小值，方便取最大值。

状态转移:
注意 k 的限制，在选择 buy 的时候相当于开启了一次交易，那么对于昨天来说，交易次数的上限 k 应该减小 1。

dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
              max( 今天选择 rest,        今天选择 sell       )
解释：今天我没有持有股票，有两种可能，我从这两种可能中求最大利润：
1. 我昨天就没有持有，且截至昨天最大交易次数限制为 k；然后我今天选择 rest，
所以我今天还是没有持有，最大交易次数限制依然为 k。
2. 我昨天持有股票，且截至昨天最大交易次数限制为 k；但是今天我 sell 了，
所以我今天没有持有股票了，最大交易次数限制依然为 k。

dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
              max( 今天选择 rest,         今天选择 buy         )
解释：今天我持有着股票，最大交易次数限制为 k，那么对于昨天来说，有两种可能，
我从这两种可能中求最大利润
1. 我昨天就持有着股票，且截至昨天最大交易次数限制为 k；然后今天选择 rest，
所以我今天还持有着股票，最大交易次数限制依然为 k。
2. 我昨天本没有持有，且截至昨天最大交易次数限制为 k - 1；但今天我选择 buy，
所以今天我就持有股票了，最大交易次数限制为 k。

*/
func maxProfit1(k int, prices []int) int {
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
	k = k + 1 //题目k给的是最大交易次数,所以数组长度必须为k+1才能保证最大交易次数限制为k

	/*
		dp[i][j][0]:表示第i天,最大交易次数为j次,不持有股票时的最大利润
		dp[i][j][1]:表示第i天,最大交易次数为j次,持有股票时的最大利润
	*/
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

/*
解法2 解法1的状态压缩
dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
1. 根据状态转移方程可以发现如果不关心dp三维数组的第2和第3维,
其实dp[i]只和dp[i-1]相关.所以可以考虑将第1维优化掉
2. 第3维只有0和1两种状态也可以考虑把第3维和第2维互换,
然后用2个变量优化掉
最后只用2个长度为k+1的1维数组表示 dp
3. 本题中的k最大值可以达到10^9,因为n天最多只能进行n/2次交易
因此可以取k和n/2中的较小值作为k然后再进行动态规划

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
	var min = func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	n := len(prices)
	k = min(k, n>>1) + 1 //题目k给的是最大交易次数,所以数组长度必须为k+1才能保证最大交易次数限制为k

	dp0 := make([]int, k) //不持有股票时的最大利润
	dp1 := make([]int, k) //持有股票时的最大利润

	for j := 0; j < k; j++ {
		dp1[j] = -prices[0]
	}
	dp0[0] = 0 //代替 dp[i-1][j-1][0], 初始化时相当于 dp[0][0][0]
	dp0[1] = 0 //代替 dp[i-1][j][0], 初始化时相当于 dp[0][1][0]

	for i := 1; i < n; i++ {
		for j := 1; j < k; j++ {
			dp0[j] = max(dp0[j], dp1[j]+prices[i])
			dp1[j] = max(dp1[j], dp0[j-1]-prices[i])
		}
	}

	return dp0[k-1]
}

// @lc code=end
