package main

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start

/*
解法1: 动态规划
状态定义:
dp[i] 表示凑齐金额 i 所需的最少硬币总数

状态初始值
dp[0]=0 表示凑够金额0需要0个硬币

转移方程
如果当前硬币面值为 coin, 如果用到该种硬币,则所需最少硬币为
dp[i]=dp[i-coin]+1
用1枚当前硬币后,剩余需要凑够的钱为 i-coin
*/
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

/*
解法2: 暴力递归
超时,只为拓展解题思路
此外还可以用利用hashMap记忆
*/
func coinChange2(coins []int, amount int) int {
	if len(coins) < 1 {
		return -1
	}
	if amount < 1 {
		return 0
	}

	//凑够leftAmount所需的最少硬币总数
	var change func(leftAmount int) int
	change = func(leftAmount int) int {
		if leftAmount < 0 {
			return -1
		}
		if leftAmount == 0 {
			return 0
		}
		for _, v := range coins {
			if v == leftAmount {
				return 1
			}
		}

		min := -1 //初始化为无法凑出
		for _, v := range coins {
			cur := change(leftAmount - v)
			if cur == -1 {
				//如果用到该种面值的硬币会导致无法凑出,直接跳过
				continue
			}
			//初次凑出或者遇到更优解就更新最优解
			if cur+1 < min || min == -1 {
				min = cur + 1
			}
		}
		return min
	}

	return change(amount)
}

// @lc code=end
