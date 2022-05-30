package main

/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start

/*
解法1 动态规划
[官方题解解法1的写法](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-iii-by-wrnt/)

状态定义:
在每天结束时,会处于以下5种状态中的其中一种
1. 未进行过任何操作；
2. 只进行过一次买操作；
3. 进行了一次买操作和一次卖操作，即完成了一笔交易；
4. 在完成了一笔交易的前提下，进行了第二次买操作；
5. 完成了全部两笔交易。
第一种状态没利润.所以可以用4个变量分别记录后4种状态

basecase:

状态转移:
关键要理解动态规划的每一个状态,在进行转移时,
当前状态的定义不会改变,改变的只是该种状态下的值

*/
func maxProfit1(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	buy1 := -prices[0] // 只进行过一次买操作状态的最高利润
	seller1 := 0       //进行了一次买操作和一次卖操作，即完成了一笔交易状态的最高利润
	buy2 := -prices[0] //在完成了一笔交易的前提下，进行了第二次买操作状态的最高利润
	seller2 := 0       //成了全部两笔交易状态的最高利润

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		seller1 = max(seller1, buy1+prices[i])
		buy2 = max(buy2, seller1-prices[i])
		seller2 = max(seller2, buy2+prices[i])
	}

	return seller2
}

/*
解法2 动态规划(通用解法)
*/

func maxProfit2(prices []int) int {
	n := len(prices)

	/*
		dp[i][j][0]表示第i天今天结束时最多进行j笔交易,不持有股票时能获得的最大利润
		dp[i][j][1]表示第i天今天结束时最多进行j笔交易,持有股票时能获得的最大利润
	*/
	dp := make([][3][2]int, n)
	for i := 0; i < 3; i++ {
		dp[0][i][0] = 0          // 第0天时结束时,如果不持有股票,不管如果限制交易次数,能获得的最大利润就是0
		dp[0][i][1] = -prices[0] // 第0天时结束时,不管如何限制交易次数,因为最后手里要有股票,所以最后必定要亏-prices[0]
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < n; i++ {
		for j := 1; j < 3; j++ {
			/*
			   1. 昨天持有,今天卖掉
			   2. 昨天不持有,今天也不持有
			   3. 昨天持有, 今天卖掉, 今天买后卖掉 (等价于1)
			   4. 昨天不持有, 今天买后卖掉 (等价于2)
			*/
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])

			/*
			   1. 昨天持有,今天不操作
			   2. 昨天不持有,今天买
			   3. 昨天持有,今天卖掉后又买入(等价于1)
			   4. 昨天不持有,今天买后卖,又买(等价于2)
			*/
			dp[i][j][1] = max(dp[i-1][j-1][0]-prices[i], dp[i-1][j][1])
		}
	}
	// fmt.Printf("dp=%+v\n", dp)
	return dp[n-1][2][0]

}

/*
解法3 解法2的状态压缩
*/
func maxProfit(prices []int) int {
	n := len(prices)

	/*
		dp[i][j][0]表示第i天今天结束时最多进行j笔交易,不持有股票时能获得的最大利润
		dp[i][j][1]表示第i天今天结束时最多进行j笔交易,持有股票时能获得的最大利润
	*/
	dp := [3][2]int{}
	for i := 0; i < 3; i++ {
		dp[i][0] = 0          // 第0天时结束时,如果不持有股票,不管如果限制交易次数,能获得的最大利润就是0
		dp[i][1] = -prices[0] // 第0天时结束时,不管如何限制交易次数,因为最后手里要有股票,所以最后必定要亏-prices[0]
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < n; i++ {
		for j := 1; j < 3; j++ {
			/*
			   1. 昨天持有,今天卖掉
			   2. 昨天不持有,今天也不持有
			   3. 昨天持有, 今天卖掉, 今天买后卖掉 (等价于1)
			   4. 昨天不持有, 今天买后卖掉 (等价于2)
			*/
			dp[j][0] = max(dp[j][0], dp[j][1]+prices[i])

			/*
			   1. 昨天持有,今天不操作
			   2. 昨天不持有,今天买
			   3. 昨天持有,今天卖掉后又买入(等价于1)
			   4. 昨天不持有,今天买后卖,又买(等价于2)
			*/
			dp[j][1] = max(dp[j-1][0]-prices[i], dp[j][1])
		}
	}
	// fmt.Printf("dp=%+v\n", dp)
	return dp[2][0]

}

// @lc code=end
