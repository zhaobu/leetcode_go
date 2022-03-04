package main

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start

/*
解法1 动态规划
dp[i] = dp[i-1] + dp[i-2]
*/
func climbStairs1(n int) int {
	if n < 3 {
		return n
	}
	//dp[i]表示爬到i+1阶楼梯的方法总数
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}

/*
解法2 动态规划 优化空间复杂度
因为 dp[i] = dp[i-1] + dp[i-2]
所以可以只用3个变量来保存
*/
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	pre2 := 1 // 1个台阶
	pre1 := 2 //2个台阶
	cur := 0  //n级台阶
	for i := 3; i <= n; i++ {
		cur = pre1 + pre2 //dp[i] = dp[i-1] + dp[i-2]
		pre2 = pre1       //dp[i-2]的值为dp[i-1]
		pre1 = cur        //更新dp[i-1]的值为dp[i]
	}
	return cur
}

// @lc code=end
