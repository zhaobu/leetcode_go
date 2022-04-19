package main

/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start

/*
解法1 动态规划
*/
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	if m < 1 || n < 1 {
		return 0
	}
	//dp[i][j]表示以matrix[i][j]为右下方的只包含'1'的最大正方形的边长
	dp := make([][]int, m)

	//初始化列
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	maxSide := 0 //记录的最长边长
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if i > 0 && j > 0 {
				if dp[i-1][j-1] == 0 { //左上为0时,最大边长只能为1
					dp[i][j] = 1
				} else {
					//上边可能最长的边
					up := min(dp[i-1][j], dp[i-1][j-1])
					//左边可能最长的边
					left := min(dp[i][j-1], dp[i-1][j-1])
					//最终的边长最小值为min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
					dp[i][j] = min(up, left) + 1
				}
			} else { //第0行或者第0列
				dp[i][j] = 1
			}
			if dp[i][j] > maxSide {
				maxSide = dp[i][j]
			}
		}
	}

	return maxSide * maxSide
}

// @lc code=end
