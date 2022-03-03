package main

/*
 * [面试题 08.11. 硬币](https://leetcode-cn.com/problems/coin-lcci/)
 */

/*
 解法1: 动态规划,完全不优化版,timeout
 背包九讲里面的完全背包类似.
 dp定义: dp[i][j]表示前i种硬币凑齐j分的方法总数
 basecase:
如果待凑的是0分,则只有一个解就是所有硬币个数都为0  dp[i][0]=1
如果硬币种类数为0, 则凑齐任意分的方法总数都为0 dp[0][i]=0
*/
func waysToChange(n int) int {
	if n == 0 {
		return 1
	}
	coins := []int{1, 5, 10, 25}
	//dp[i][j]表示前i种硬币凑齐j分的方法总数
	dp := make([][]int, len(coins)+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= n; j++ {
			// fmt.Printf("i=%d\n", i )
			for k := 0; k*coins[i-1] <= j; k++ {
				dp[i][j] += dp[i-1][j-k*coins[i-1]]
			}
		}
	}
	return dp[len(coins)][n] % 1000000007
}
