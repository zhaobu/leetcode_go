package main

import "math/big"

/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start

/*
解法1 dfs暴力(超时)
*/
func uniquePaths1(m int, n int) int {
	ret := 0
	m, n = m-1, n-1
	dirs := [][2]int{{1, 0}, {0, 1}}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i == m && j == n {
			ret++
			return
		}

		for _, v := range dirs {
			i1, j1 := i+v[0], j+v[1]
			if i1 <= m && j1 <= n {
				dfs(i1, j1)
			}
		}
	}

	dfs(0, 0)

	return ret
}

/*
解法2 动态规划
*/
func uniquePaths2(m int, n int) int {

	/*
		dp[i][j]: 表示到达(i,j)网格的不同路径数量
	*/
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	return dp[m-1][n-1]
}

/*
解法3 动态规划(优化空间复杂度)
*/
func uniquePaths3(m int, n int) int {

	/*
		dp[i][j]: 表示到达(i,j)网格的不同路径数量
	*/
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[j] = 1
			} else {
				dp[j] = dp[j-1] + dp[j]
			}
		}
	}

	return dp[n-1]
}

/*
解法4 组合数学
*/
func uniquePaths(m, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}

// @lc code=end
