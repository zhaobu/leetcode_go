package main

import "math"

/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start

/*
解法1 动态规划 股票问题通用解法
参考: https://labuladong.gitee.io/algo/3/27/100/

状态定义:
dp[i][0]表示第i天,不持有股票时的最大利润
dp[i][1]表示第i天,持有股票时的最大利润

basecase:
dp[0][0] = 0
dp[0][1] = -prices[0]
dp[1][0] = max(dp[0][0], dp[0][1]+prices[1])
dp[1][1] = max(dp[0][1], -prices[1])
1. 第0天买,第1天不操作 2. 第1天买股票,不存在第0天买了之后卖,因为不允许

状态转移方程:
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])
解释：第 i 天选择 buy 的时候，要从 i-2 的状态转移，而不是 i-1 。

*/
func maxProfit1(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[1][0] = max(dp[0][0], dp[0][1]+prices[1])
	//1. 第0天买,第1天不操作 2. 第1天买股票,不存在第0天买了之后卖,因为不允许
	dp[1][1] = max(dp[0][1], -prices[1])

	for i := 2; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
	}

	return dp[len(prices)-1][0]
}

/*
解法2 动态规划
解法1状态压缩版
根据解法1的basecase和状态转移方程观察:
1. 只有3个变量 dp[i-1][0], dp[i-1][1], dp[i-2][0] 三个变量在更新
2. 在计算完第i次后 dp[i][0] 其实是第 i+1 次循环时的 dp[i-1][0]
dp[i][1] 其实是第 i+1 次循环时的 dp[i-1][1]
所以可以直接滚动更新即可
3. 第 i+1 次循环的 dp[i-2][0] 和第i次循环的 dp[i-1][0],所以要提前保存该值

*/

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dpi10 := 0
	dpi11 := -prices[0]
	dpi20 := 0

	for i := 1; i < len(prices); i++ {
		olddpi10 := dpi10 //必须要提前保存,不然更新后再保存,在下一轮循环时对应不上
		dpi10 = max(dpi10, dpi11+prices[i])
		dpi11 = max(dpi11, dpi20-prices[i])
		dpi20 = olddpi10
	}

	return dpi10
}

/*
解法3 动态规划
[根据英文网站高赞回答实现] (https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/discuss/75928/Share-my-DP-solution-(By-State-Machine-Thinking))

状态定义:
s0:
定义: 未持有股票,卖股票后下下一天:
from: 1. reset,也就是今天啥都没干 2. 由s2转移而来
to: s1 非冷冻期后可以买股票

s1:
定义: 持有股票:
from: 1. reset,也就是昨天买了股票,今天啥都没干 2. 昨天没买也没卖,今天买了股票, 由s0转移而来
to: s2 把持有的股票卖了后就进入冷冻期

s2:
定义: 未持有股票,卖股票后下一天,也就是冷冻期
from: s1 由s1转移而来
to: s0 转向冷冻期的下一天

*/
func maxProfit3(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	s0 := make([]int, len(prices))
	s1 := make([]int, len(prices))
	s2 := make([]int, len(prices))

	s0[0] = 0           // 未持有股票
	s1[0] = -prices[0]  //持有股票
	s2[0] = math.MinInt //冷冻期,未定义状态,但要求的是最大值,所以初始化为最小值

	for i := 1; i < len(prices); i++ {
		s0[i] = max(s0[i-1], s2[i-1])           //1. 从昨天的s0 reset 2. 从昨天的s2转移
		s1[i] = max(s1[i-1], s0[i-1]-prices[i]) // 1. 从昨天的s1 reset 2. 从昨天的s0朱阿姨
		s2[i] = s1[i-1] + prices[i]             // 从昨天的s1转移
	}

	return max(s0[len(prices)-1], s2[len(prices)-1])
}

/*
解法4 解法3的状态压缩
*/
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	s0 := 0           // 未持有股票
	s1 := -prices[0]  //持有股票
	s2 := math.MinInt //冷冻期,未定义状态,但要求的是最大值,所以初始化为最小值

	for i := 1; i < len(prices); i++ {
		olds0, olds1 := s0, s1        //必须提前保存,同一轮循环中使用同一个值
		s0 = max(s0, s2)              //1. 从昨天的s0 reset 2. 从昨天的s2转移
		s1 = max(s1, olds0-prices[i]) // 1. 从昨天的s1 reset 2. 从昨天的s0转移
		s2 = olds1 + prices[i]        // 从昨天的s1转移
	}

	return max(s0, s2)
}

// @lc code=end
