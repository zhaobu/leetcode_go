package main

import "math"

/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start

/*
解法1 动态规划

dp[i] 表示将正整数 i 拆分成至少两个正整数的和之后，这些正整数的最大乘积。
0 不是正整数，1 是最小的正整数，0 和 1 都不能拆分，因此dp[0]=dp[1]=0。

当 i>=2 时,假设 j 是拆分出来的第1个正整数, j的取值范围是[1,i).则有2种方法
1. 分为2部分,第一个数是j,第二个数是 i-j. 也就是队剩下的i-j不继续拆分. 此时乘积是 j*(i-j)
2. 分为多部分,第一个数是j,剩下的 i-j,再进行拆分成至少2个数, 此时的乘积是 j*dp[i-j]
*/

func integerBreak1(n int) int {
	if n <= 2 {
		return 1
	}
	/*
		1. dp[i] 表示将正整数 i 拆分成至少两个正整数的和之后，这些正整数的最大乘积。
		2. 0 不是正整数，1 是最小的正整数，0 和 1 都不能拆分，因此dp[0]=dp[1]=0。
	*/
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	dp[2] = 1

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 3; i <= n; i++ {
		for j := 1; j <= (i >> 1); j++ {
			dp[i] = max(max(j*dp[i-j], j*(i-j)), dp[i])
		}
	}

	return dp[n]
}

/*
解法2 官方解法2
https://leetcode.cn/problems/integer-break/solution/zheng-shu-chai-fen-by-leetcode-solution/
*/
func integerBreak2(n int) int {
	if n < 4 {
		return n - 1
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = max(max(2*(i-2), 2*dp[i-2]), max(3*(i-3), 3*dp[i-3]))
	}
	return dp[n]

}

/*
解法3 数学
*/
func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}
	quotient := n / 3
	remainder := n % 3
	if remainder == 0 {
		return int(math.Pow(3, float64(quotient)))
	} else if remainder == 1 {
		return int(math.Pow(3, float64(quotient-1))) * 4
	}
	return int(math.Pow(3, float64(quotient))) * 2
}

// @lc code=end
