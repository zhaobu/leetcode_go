package main

/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start

/*
解法1 动态规划
*/
func change1(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}

	/*
		dp[i]表示凑齐金额i的不同coins组合种类数
	*/
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, v := range coins {
		for i := v; i <= amount; i++ {
			if dp[i-v] > 0 {
				dp[i] += dp[i-v]
			}
		}
	}
	// fmt.Printf("dp=%+v\n", dp)
	return dp[amount]
}

/*
解法2 动态规划
状态定义:
状态有两个，就是「背包的容量」和「可选择的物品」，选择就是「装进背包」或者「不装进背包」
dp[i][j] 表示只使用前 i 个物品（可以重复使用），当背包容量为 j 时，
有 dp[i][j] 种方法可以装满背包.
所以题目要求的就是dp[len(coins)][amount]

basecase:
dp[0][j]=0 不使用硬币,能凑出金额j的方法为0,(dp[0][1]=1)
dp[i][0]=1 凑齐金额0的方式, 就是不使用任何硬币这一种方式

转移方程:
不使用当前硬币:
dp[i][j] = dp[i-1][j]
使用当前硬币
dp[i][j] = dp[i-1][j]+dp[i][j-coins[i-1]]
(由于定义中的 i 是从 1 开始计数的，所以 coins 的索引是 i-1 时表示第 i 个硬币的面值)
*/
func change2(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}

	/*
		dp[i]表示凑齐金额i的不同coins组合种类数
	*/
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}
	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i-1][j] // 不使用当前硬币
			coin := coins[i-1]
			if j >= coin { //使用当前硬币
				dp[i][j] += dp[i][j-coin]
			}
		}
	}
	return dp[len(coins)][amount]
}

/*
解法3 动态规划
解法2的空间压缩
*/
func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}

	/*
		dp[i]表示凑齐金额i的不同coins组合种类数
	*/
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

// @lc code=end
