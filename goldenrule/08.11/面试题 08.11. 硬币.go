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
func waysToChange1(n int) int {
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

/*
解法2 动态规划优化,也就是官方解法1,对解法的优化
https://leetcode-cn.com/problems/coin-lcci/solution/ying-bi-by-leetcode-solution/
f(i,v)=f(i−1,v)+f(i−1,v−ci)+f(i−1,v−2ci​)⋯f(i−1,v−kci​)
主要是根据
dp[i,j] = dp[i-1][j] + dp[i-1][j-coins[i]] + dp[i-1][j-2coins[i]] + ...  +dp[i-1][j-kcoins[i]]
dp[i][j-coins[i]] = dp[i-1][j-coins[i]] + dp[i-1][j-2coins[i]] + ...  +dp[i-1][j-kcoins[i]]
2个递推公式得出优化后的递推公式:
dp[i][j] = dp[i-1][j] + dp[i][k-coins[i]]
其中 k=j/coins[i]

*/

func waysToChange2(n int) int {
	if n == 0 {
		return 1
	}
	var coins = []int{1, 5, 10, 25}

	//dp[i][j]表示前i种硬币凑齐j分的方法总数
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= n; j++ {
			dp[j] = dp[j] + dp[j-coins[i]]
		}
	}
	return dp[n] % 1000000007
}

/*
解法3 官方题解解法2
数学公式推导
https://leetcode-cn.com/problems/coin-lcci/solution/ying-bi-by-leetcode-solution/
*/
func waysToChange(n int) int {
	ans := 0
	mod := 1000000007
	for i := 0; i*25 <= n; i++ {
		rest := n - i*25
		a, b := rest/10, rest%10/5
		ans = (ans + (a+1)*(a+b+1)%mod) % mod
	}
	return ans
}
