package main

import "math"

/*
 *
 *
 * 剑指 Offer 14- I. 剪绳子
 */

// @lc code=start

/*
 解法1 动态规划

*/
func cuttingRope(n int) int {
	if n <= 2 {
		return 1
	}
	/*
		1. dp[i] 表示将绳子 i 拆分成至少两个正整数的和之后，这些正整数的最大乘积。
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

func cuttingRope2(n int) int {
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
