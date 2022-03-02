package main

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	if len(coins) < 1 {
		return -1
	}
	if amount < 1 {
		return 0
	}
	//dp[i] 表示凑齐金额 i 所需的最少硬币总数
	dp := make([]int, amount+1)

	//从1开始到amount,计算每种金额需要的最少硬币总数
	for i := 1; i < len(dp); i++ {
		dp[i] = -1 //默认初始化为找不出
		for _, v := range coins {
			if i < v || dp[i-v] == -1 {
				/*
					1. 如果金币面值超过金额,跳过
					2. 根据dp记录知道如果用了该面值的硬币后凑不出,跳过
				*/
				continue
			}
			cur := dp[i-v] + 1 //用1枚当前面值硬币 + 剩下金额需要的最少硬币数
			if cur < dp[i] || dp[i] == -1 {
				/*
					1. 如果是dp[i]=-1,表明是第一次找到凑出硬币的方法,记录所需硬币总数
					2. 如果所需硬币总数更小,就更新记录
				*/
				dp[i] = cur
			}
		}
	}
	return dp[amount]
}

// @lc code=end
