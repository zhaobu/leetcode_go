package main

/*
 *
 *
 * 剑指 Offer 10- II. 青蛙跳台阶问题
 */

// @lc code=start

/*
 解法1
*/
func numWays(n int) int {
	if n < 2 {
		return 1
	}
	dp0, dp1 := 1, 2
	for i := 3; i <= n; i++ {
		tmp := (dp1 + dp0) % 1000000007
		dp0 = dp1
		dp1 = tmp
	}

	return dp1
}

// @lc code=end
