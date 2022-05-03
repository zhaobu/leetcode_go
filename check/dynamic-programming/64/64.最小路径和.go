package main

/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start

/*
解法1 动态规划
*/
func minPathSum1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = grid[i][0]
		if i > 0 {
			dp[i][0] += dp[i-1][0]
		}
	}

	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

/*
解法2 动态规划(解法1空间复杂度优化)
1. 观察转移方程dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
dp[i][j]只和dp[i-1][j]和dp[i][j-1]也就是本行前一列dp和上一行同一列dp有关.
2. 考虑到如果从左往右,从上往下遍历:
dp[j-1]就是本行前一列dp, dp[j]就是上一行同一列dp.
3. 二维dp可以空间压缩为1维dp
*/
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i != 0 {
				if j > 0 {
					dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
				} else {
					dp[j] += grid[i][j]
				}
			} else {
				dp[j] = grid[i][j]
				if j > 0 {
					dp[j] += dp[j-1]
				}
			}
		}
	}
	return dp[n-1]
}

/*
解法3 动态规划(解法2的更简洁写法)
*/
func minPathSum3(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[j] += grid[i][j]
			if j > 0 {
				if i != 0 {
					dp[j] = min(dp[j-1]+grid[i][j], dp[j])
				} else {
					dp[j] += dp[j-1]
				}
			}
		}
	}
	return dp[n-1]
}

// @lc code=end
